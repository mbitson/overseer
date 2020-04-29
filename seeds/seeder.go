package seeds

import (
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
	pk "github.com/mbitson/overseer/utilities/pbkdf2"
)

func Seed() {
	isSeeded := isSeeded()
	if isSeeded == false {
		seedUsers()
		seedApplicationTypes()
		seedMonitorTypes()
		seedSitesAndMonitors()
	}
}

func seedUsers() {
	first := beego.AppConfig.String("admin_first")
	last := beego.AppConfig.String("admin_last")
	email := beego.AppConfig.String("admin_email")
	user := models.AuthUser{First: first, Last: last, Email: email}
	h := pk.HashPassword(beego.AppConfig.String("admin_init_pw"))
	user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&user)

	logEntry1 := models.Log{Action: "User.Register", Ip: "127.0.0.1", User_id: 1, Action_id: 1}
	o.Insert(&logEntry1)
}

func seedSitesAndMonitors() {
	domains := beego.AppConfig.Strings("monitored_sites")
	for _, domain := range domains {
		var site *models.Site
		site = seedSite(domain)
		seedMonitor(site)
		seedAlert(site)
	}
}

func seedAlert(site *models.Site) {
	o := orm.NewOrm()
	o.Using("default")

	name := "Default Alert"
	email := beego.AppConfig.String("admin_email")
	sms := beego.AppConfig.String("admin_sms")
	urgency := 4
	Alert := models.Alert{Name: name, Site: site, Email: email, Sms: sms, Urgency: urgency}
	o.Insert(&Alert)

	return
}

func seedMonitor(site *models.Site) {
	o := orm.NewOrm()
	o.Using("default")

	interval := 120
	type_id := 1
	monitorType := models.MonitorType{Id: int(type_id)}

	monitor := models.Monitor{Name: "Default", Site: site, Interval: interval, Monitor_type: &monitorType}
	m_id, err := o.Insert(&monitor)
	if err != nil {
		//fmt.Println(err)
	}
	monitor_id := int(m_id)

	logEntry2 := models.Log{Action: "Monitor.Create", Ip: "127.0.0.1", User_id: 1, Action_id: monitor_id}

	_, err2 := o.Insert(&logEntry2)
	if err2 != nil {
		//fmt.Println(err2)
	}

	return
}

func seedSite(domain string) *models.Site {
	user_id := 1
	site := models.Site{User: &models.AuthUser{Id: user_id}, Domain: domain}
	o := orm.NewOrm()
	o.Using("default")
	id, _ := o.Insert(&site)
	site_id := int(id)

	logEntry1 := models.Log{Action: "Site.Create", Ip: "127.0.0.1", User_id: 1, Action_id: site_id}
	o.Insert(&logEntry1)
	site.Id = site_id

	return &site
}

func seedMonitorTypes() {
	o := orm.NewOrm()
	o.Using("default")

	DomainMonitor := models.MonitorType{
		Name:        "Domain",
		Description: "A simple domain monitor that will connect to a domain or IP and report back the ping, page load time, and server response time.",
	}
	o.Insert(&DomainMonitor)
}

func seedApplicationTypes() {
	o := orm.NewOrm()
	o.Using("default")

	wordpressApp := models.ApplicationType{
		Name:        "WordPress",
		Description: "WordPress is a CMS designed for publishing a blog.",
		Website:     "http://www.wordpress.org/",
	}
	o.Insert(&wordpressApp)

	magentoApp := models.ApplicationType{
		Name:        "Magento",
		Description: "Magento is an eCommerce application.",
		Website:     "https://magento.com/",
	}
	o.Insert(&magentoApp)

	magento2App := models.ApplicationType{
		Name:        "Magento 2",
		Description: "Magento 2 is an eCommerce application.",
		Website:     "https://magento.com/",
	}
	o.Insert(&magento2App)

	genericApp := models.ApplicationType{
		Name:        "Generic",
		Description: "A generic monitor which checks for HTTP response codes only.",
		Website:     "https://mbitson.com/",
	}
	o.Insert(&genericApp)
}

func isSeeded() bool {
	o := orm.NewOrm()
	o.Using("default")

	domainMonitorSearch := models.MonitorType{Name: "Domain"}
	domainMonitorQuery := o.Read(&domainMonitorSearch, "Name")

	if domainMonitorQuery == orm.ErrNoRows {
		return false
	} else if domainMonitorQuery == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}
