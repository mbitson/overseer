package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mbitson/overseer/routers"
	"github.com/mbitson/overseer/seeds"
	_ "github.com/mbitson/overseer/tasks"
)

func main() {
	// Configure Application-level Options
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"

	// Registered MySQL as the Preferred Driver
	regDriverErr := orm.RegisterDriver("mysql", orm.DRMySQL)
	if regDriverErr != nil {
		panic(regDriverErr)
	}

	// Connect specifically to the app.conf connection string
	dbUrl := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") +
		"@" + beego.AppConfig.String("mysqlurls") + "/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	regDbErr := orm.RegisterDataBase("default", "mysql", dbUrl)
	if regDbErr != nil {
		panic(regDbErr)
	}

	// Migrate (Create/update any database tables as necessary)
	syncErr := orm.RunSyncdb("default", false, false)
	if syncErr != nil {
		fmt.Println(syncErr)
	}

	// Seed (Creates any database records necessary)
	seeds.Seed()

	// Runs the server to respond to requests
	beego.Run()
}
