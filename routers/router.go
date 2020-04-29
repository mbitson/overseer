package routers

import (
	"github.com/astaxie/beego"
	"github.com/mbitson/overseer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	beego.Router("/user/login/", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/login/:back", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/user/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/user/profile", &controllers.UserController{}, "get,post:Profile")
	beego.Router("/user/verify/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.UserController{}, "get:Verify")
	beego.Router("/user/remove", &controllers.UserController{}, "get,post:Remove")
	beego.Router("/notice", &controllers.UserController{}, "get:Notice")
	beego.RESTRouter("/api/sites", &controllers.SitesApiController{})
	beego.RESTRouter("/api/monitorRuns", &controllers.MonitorRunsApiController{})
	beego.RESTRouter("/api/monitors", &controllers.MonitorsApiController{})
	beego.RESTRouter("/api/alerts", &controllers.AlertsApiController{})
	beego.RESTRouter("/api/applicationTypes", &controllers.ApplicationTypesApiController{})
	beego.RESTRouter("/api/contacts", &controllers.ContactsApiController{})
	beego.RESTRouter("/api/groups", &controllers.GroupsApiController{})
	beego.Router("/api/groups/saveContacts", &controllers.GroupsApiController{}, "post:SaveContacts")
	beego.Router("/api/live/join", &controllers.LiveController{}, "get:Join")
	beego.Router("/api/alerts/recent/:id", &controllers.AlertsApiController{}, "get:GetRecent")
	beego.Router("/api/monitorRuns/:id/:start/:end", &controllers.MonitorRunsApiController{}, "get:GetOne")
	beego.Router("/api/events", &controllers.EventsApiController{}, "get:GetAll")
}
