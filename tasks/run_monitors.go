package tasks

import (
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	live "github.com/mbitson/overseer/controllers"
	"github.com/mbitson/overseer/models"
	"github.com/parnurzeal/gorequest"
	"github.com/robfig/cron"
	"gopkg.in/gomail.v1"
	"net"
	"time"
)

func init() {
	c := cron.New()
	err := c.AddFunc("@every 60s", run_monitors)
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
}

func run_monitors() {
	fmt.Print("\r\nScanning site statuses. \r\n")
	runTime := time.Now()

	var monitors []models.Monitor
	o := orm.NewOrm()
	o.QueryTable("monitor").RelatedSel("site").All(&monitors)

	fmt.Print("Processing runs... \r\n")
	runs := get_runs(monitors, runTime)

	for _, thisRun := range runs {
		// Insert this run
		o.Insert(thisRun)
		o.Read(thisRun.Monitor)
		o.Read(thisRun.Monitor.Site)
		o.Read(thisRun.Monitor.Site.User)

		event := live.NewEvent("MONITOR.RUN", thisRun.Monitor.Site.User.Email, []live.EventData{thisRun})
		live.BroadcastEvent(event)

		// Update this monitor's checked date
		thisRun.Monitor.Checked_date = runTime
		o.Update(thisRun.Monitor, "Checked_date")

		// Update site to reflect status code returned.
		site := thisRun.Monitor.Site
		if thisRun.Status_code == 200 {
			site.Status = 1
		} else {
			site.Status = 0
		}
		o.Update(site, "Status")

		fmt.Print(site.Status)

		// Check all alerts, trigger necessary notifications
		check_alerts(thisRun)
	}

	return
}

func get_runs(monitors []models.Monitor, runTime time.Time) []*models.MonitorRun {
	runChannel := make(chan *models.MonitorRun, len(monitors)) // buffered, max length same as number of monitors

	for _, monitor := range monitors {
		go check_monitor(monitor, runTime, runChannel)
	}

	var runs []*models.MonitorRun

	for {
		select {
		case run := <-runChannel:
			runs = append(runs, run)
			if len(runs) == len(monitors) {
				return runs
			}
		case <-time.After(60 * time.Second):
		}
	}
	return runs
}

func check_monitor(monitor models.Monitor, runTime time.Time, runChannel chan *models.MonitorRun) {
	// Firstly, check that this monitor needs to fire this round. See last fired date and interval and compare to now.
	interval := float64(monitor.Interval)
	timeSince := time.Since(monitor.Checked_date.Add(-4 * time.Hour)).Seconds()
	if interval < timeSince {
		// Get URL from Domain
		url := fmt.Sprint("http://", monitor.Site.Domain)

		client := gorequest.New().Timeout(2 * time.Millisecond)
		pageStart := time.Now()
		page, _, err := client.Get(url).Timeout(60 * time.Second).End()
		var statusCode int
		statusCode = 0
		if err == nil {
			statusCode = page.StatusCode
		} else {
			statusCode = 500
		}
		pageElapsed := time.Since(pageStart).Nanoseconds() / 1000000

		ipStart := time.Now()
		net.LookupIP(monitor.Site.Domain)
		ipElapsed := time.Since(ipStart).Nanoseconds() / 1000000

		var monitorRun models.MonitorRun
		monitorRun.Monitor = &monitor
		monitorRun.Ping = ipElapsed
		monitorRun.Response_time = pageElapsed
		monitorRun.Status_code = statusCode
		monitorRun.Time_run = runTime

		runChannel <- &monitorRun
	}
	return
}

func check_alerts(monitorRun *models.MonitorRun) {
	o := orm.NewOrm()
	o.Using("default")

	siteAlerts := get_site_alerts(monitorRun.Monitor.Site)
	monitorAlerts := get_monitor_alerts(monitorRun.Monitor)
	alerts := append(siteAlerts, monitorAlerts...)

	for _, alert := range alerts {
		alertFlag := false
		if alert.Urgency > 1 {
			alertFlag = true
			var monitorRuns []*models.MonitorRun
			o.QueryTable("monitor_run").Filter("Monitor", monitorRun.Monitor).Limit(alert.Urgency).OrderBy("-time_run").All(&monitorRuns)
			for _, monitorRunHist := range monitorRuns {
				if monitorRunHist.Status_code == 200 {
					alertFlag = false
				}
			}
		} else {
			if monitorRun.Status_code != 200 {
				alertFlag = true
			}
		}

		if alertFlag == true {
			// Email Alert
			email_alert(alert, monitorRun)

			// TODO: Implement SMS alerts. MessageInAction

			// Update alert's last checked date
			alert.Checked_date = time.Now()
			o.Update(alert, "Checked_date")
		} else {

		}

	}

	return
}

func email_alert(alert *models.Alert, run *models.MonitorRun) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "me@mbitson.com")
	msg.SetHeader("To", alert.Email)
	msg.SetHeader("Subject", "Site Down!", run.Monitor.Site.Domain, " - triggered by ", alert.Name)
	const timeFormat = "Jan 2, 2006 at 3:04pm (MST)"
	body := fmt.Sprint("Check yo alerts!! Alert ", alert.Name, " triggered at ", alert.Checked_date.Format(timeFormat), ". This alert was triggered with run ID ", run.Id, " returning ", run.Status_code)
	msg.SetBody("text/html", body)
	mailer := gomail.NewMailer(
		"mbitson.com",
		"go@mbitson.com",
		"~paran01d~",
		465, gomail.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)
	if err := mailer.Send(msg); err != nil {
		spew.Dump(err)
	}
}

func get_monitor_alerts(monitor *models.Monitor) []*models.Alert {
	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var alerts []*models.Alert
	o.QueryTable("alert").Filter("Monitor", monitor).All(&alerts)
	return alerts
}

func get_site_alerts(site *models.Site) []*models.Alert {
	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var alerts []*models.Alert
	o.QueryTable("alert").Filter("Site", site).All(&alerts)
	return alerts
}
