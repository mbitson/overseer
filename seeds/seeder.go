package seeds

import (
	"go.mbitson.com/models"
	"github.com/astaxie/beego/orm"
	pk "go.mbitson.com/utilities/pbkdf2"
	"encoding/hex"
	//"fmt"
)

func Seed(){
	// Seed Users
	seedUsers()
	// Seed Application Types
	seedApplicationTypes()
	// Seed Monitor Types
	seedMonitorTypes()
	// Seed Sites, Monitors, And Alerts.
	seedSitesAndMonitors()
}
func seedUsers(){
	first := "Mikel"
	last := "Bitson"
	email := "test@gmail.com"
	user := models.AuthUser{First: first, Last: last, Email: email}
	h := pk.HashPassword("123456")
	user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&user)

	logEntry1 := models.Log{Action: "User.Register", Ip: "127.0.0.1", User_id: 1, Action_id: 1}
	o.Insert(&logEntry1)
}
func seedSitesAndMonitors(){
	domains := [...]string{
		"mcg.mbitson.com",
	}
	for _, domain := range domains {
		var site *models.Site
		site = seedSite(domain)
		seedMonitor(site)
		seedAlert(site)
	}
}
func seedAlert(site *models.Site){
	o := orm.NewOrm()
	o.Using("default")

	name := "Default Alert"
	email := "me@mbitson.com"
	sms := "2313296944"
	urgency := 4
	Alert := models.Alert{Name:name,Site:site,Email:email,Sms:sms,Urgency:urgency}
	o.Insert(&Alert)

	return
}
func seedMonitor(site *models.Site){
	o := orm.NewOrm()
	o.Using("default")

	interval := 120
	type_id := 1
	monitorType := models.MonitorType{Id: int(type_id)}

	monitor := models.Monitor{Name: "Default", Site: site, Interval: interval, Monitor_type: &monitorType}
	m_id, err := o.Insert(&monitor)
	if err != nil{
		//fmt.Println(err)
		//fmt.Println(monitor.Site)
	}
	monitor_id := int(m_id)

	logEntry2 := models.Log{Action: "Monitor.Create", Ip: "127.0.0.1", User_id: 1, Action_id: monitor_id}

	_, err2 := o.Insert(&logEntry2)
	if err2 != nil{
		//fmt.Println(err2)
	}

	return
}
func seedSite(domain string) (*models.Site){
	user_id := 1
	site := models.Site{User: &models.AuthUser{Id:user_id}, Domain: domain}
	o := orm.NewOrm()
	o.Using("default")
	id, _ := o.Insert(&site)
	site_id	:= int(id)

	logEntry1 := models.Log{Action: "Site.Create", Ip: "127.0.0.1", User_id: 1, Action_id: site_id}
	o.Insert(&logEntry1)
	site.Id = site_id

	return &site 
}
func seedMonitorTypes(){
	monitorName := "Domain"
	monitorDescription := "A simple domain monitor that will connect to a domain or IP and report back the ping, page load time, and server response time."
	monitorType := models.MonitorType{Name: monitorName, Description: monitorDescription}
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&monitorType)
}

func seedApplicationTypes(){
	typeName := "WordPress"
	typeDescription := "WordPress is used in millions of sites."
	typeLogo_url := "https://s.w.org/about/images/logos/wordpress-logo-stacked-rgb.png"
	typeWebsite := "http://www.wordpress.org/"
	applicationType := models.ApplicationType{Name: typeName, Description: typeDescription, Logo_url: typeLogo_url, Website: typeWebsite}
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&applicationType)
}
