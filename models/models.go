package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

/**
 * Declare Models/Tables
 */
// Log table for action history
type Log struct {
	Id        int
	Ip        string
	User_id   int `orm:"null";`
	Action    string
	Action_id int       `orm:"null";`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
}

// Auth user is a user.
type AuthUser struct {
	Id       int
	First    string
	Last     string
	Email    string `orm:"unique"`
	Password string
	Reg_key  string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

// User's contacts
type Contact struct {
	Id    int
	User  *AuthUser `orm:"rel(fk)"`
	First string
	Last  string
	Email string `orm:"null";`
	Sms   string `orm:"null";`
}

// User's groups
type Group struct {
	Id             int
	User           *AuthUser `orm:"rel(fk)"`
	Name           string
	Contacts       string     `orm:"size(100)"`
	ContactDetails *[]Contact `orm:"-"`
}

type GroupContacts struct {
	Id        int
	GroupId   int `orm:"index"`
	ContactId int
}

// Site is a site that is registered to a user.
type Site struct {
	Id      int
	User    *AuthUser `orm:"rel(fk)"`
	Domain  string    `orm:"unique"`
	Status  int       `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

// Create a model for deployment credentials.
type Deployment struct {
	Id       int
	Site     *Site `orm:"rel(fk)"`
	Name     string
	Host     string
	Port     int `orm:"default(21)"`
	Path     string
	Username string
	Password string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

// Create a model for the
type Application struct {
	Id         int
	Type       *ApplicationType `orm:"rel(fk)"`
	Site       *Site            `orm:"rel(fk)"`
	Deployment *Deployment      `orm:"rel(fk)"`
	Created    time.Time        `orm:"auto_now_add;type(datetime)"`
}

type ApplicationType struct {
	Id          int
	Name        string `orm:"unique"`
	Description string
	Website     string
}

// Create a model for the monitor type.
type Monitor struct {
	Id           int
	Site         *Site `orm:"rel(fk)"`
	Name         string
	Interval     int
	Monitor_type *MonitorType `orm:"rel(fk)"`
	Checked_date time.Time    `orm:"type(datetime);null";`
	Created      time.Time    `orm:"auto_now_add;type(datetime)"`
}

// Create a model for the types of monitors
type MonitorType struct {
	Id          int
	Name        string
	Description string
}

// Create a model for the site statistics
type MonitorRun struct {
	Id            int
	Monitor       *Monitor `orm:"rel(fk)"`
	Alert         *Alert   `orm:"rel(fk);null";`
	Ping          int64
	Page_load     int64 `orm:"null";`
	Response_time int64
	Status_code   int
	Time_run      time.Time `orm:"type(datetime)"`
}

// Create a new type for Alerts
type Alert struct {
	Id           int
	Monitor      *Monitor `orm:"rel(fk);null"`
	Site         *Site    `orm:"rel(fk)"`
	Name         string
	Email        string `orm:"null";`
	Sms          string `orm:"null";`
	Urgency      int
	Checked_date time.Time `orm:"type(datetime);null";`
}

type AlertNotification struct {
	Id       int
	Alert    *Alert    `orm:"rel(fk)"`
	Email    string    `orm:"null";`
	Sms      string    `orm:"null";`
	Notified time.Time `orm:"type(datetime)"`
}

/**
 * Load our models into the ORM.
 */
func init() {
	// Register Models
	orm.RegisterModel(
		new(AuthUser),
		new(Contact),
		new(Group),
		new(Site),
		new(Log),
		new(Monitor),
		new(MonitorType),
		new(MonitorRun),
		new(Alert),
		new(AlertNotification),
		new(ApplicationType),
		new(Application),
		new(Deployment),
	)
}
