package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Test() {
}

func (this *MainController) activeContent(view string, layout string) {
	if layout == "" {
		layout = "layout-bare"
	}

	this.LayoutSections = make(map[string]string)

	this.LayoutSections["Head"] = "common/head.tpl"
	this.LayoutSections["Header"] = "common/header.tpl"
	this.LayoutSections["Sidebar"] = "common/sidebar.tpl"
	this.LayoutSections["Footer"] = "common/footer.tpl"

	this.Layout = layout + ".tpl"
	this.TplName = view + ".tpl"

	sess := this.GetSession("os-auth")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) bareContent(view string, layout string) {

	this.LayoutSections = make(map[string]string)

	this.LayoutSections["Head"] = "common/head.tpl"
	this.LayoutSections["Header"] = "common/bare-header.tpl"
	this.LayoutSections["Footer"] = "common/bare-footer.tpl"

	this.Layout = "layout-bare.tpl"
	this.TplName = view + ".tpl"

	sess := this.GetSession("os-auth")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) ajaxContent(content interface{}) {
	this.Data["json"] = &content
	this.ServeJSON()
}

func (this *MainController) Get() {
	this.activeContent("index", "")

	//******** This page requires login
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}
}

func (this *MainController) Notice() {
	this.activeContent("notice", "")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}

func (this *MainController) logRouteAction(user int) {
	cont, act := this.GetControllerAndAction()
	s := []string{strings.Replace(cont, "Controller", "", 1), ".", act}
	action := strings.Join(s, "")
	logEntry := models.Log{Action: action, Ip: this.Ctx.Request.RemoteAddr, User_id: user, Action_id: user}

	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&logEntry)
}

func (this *MainController) logCustomAction(action string, user int) {
	logEntry := models.Log{Action: action, Ip: this.Ctx.Request.RemoteAddr, User_id: user, Action_id: user}
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&logEntry)
}

func (this *MainController) logCustomActionAndId(action string, id int, user int) {
	logEntry := models.Log{Action: action, Ip: this.Ctx.Request.RemoteAddr, User_id: user, Action_id: id}
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(&logEntry)
}
