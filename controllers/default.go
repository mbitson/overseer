package controllers

import (
	"github.com/astaxie/beego"
	"go.mbitson.com/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"bitbucket.org/tebeka/selenium"
	"fmt"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Test(){
	fmt.Println("Requesting capabilities and creating session...")
	caps := selenium.Capabilities{
		"browserName": "firefox",
		"javascriptEnabled": true,
		"takesScreenshot": true,
	}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil{
		fmt.Println(err)
	}
	defer wd.Quit()

	fmt.Println("Connecting to URL with session...")
	// Get simple playground interface
	wd.Get("http://www.grandvillages.com")


	fmt.Println("Getting Body element...")
	// Get the result
	output := ""
	for {
		output, _ = wd.PageSource()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("Setting return JSON...")
	this.Data["json"] = &output

	fmt.Println("Serve JSON.")

	this.ServeJson()
}

func (this *MainController) activeContent(view string, layout string) {
	if layout == "" {
		layout = "layout-bare"
	}

	this.LayoutSections = make(map[string]string)

	this.LayoutSections["Header"] = "common/header.tpl"
	this.LayoutSections["Sidebar"] = "common/sidebar.tpl"
	this.LayoutSections["Footer"] = "common/footer.tpl"

	this.Layout = layout + ".tpl"
	this.TplNames = view + ".tpl"

	sess := this.GetSession("go.mbitson.com")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) bareContent(view string, layout string) {

	this.LayoutSections = make(map[string]string)

	this.LayoutSections["Header"] = "common/bare-header.tpl"
	this.LayoutSections["Footer"] = "common/bare-footer.tpl"

	this.Layout = "layout-bare.tpl"
	this.TplNames = view + ".tpl"

	sess := this.GetSession("go.mbitson.com")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
	}
}

func (this *MainController) ajaxContent(content interface{}) {
	this.Data["json"] = &content
	this.ServeJson()
}

func (this *MainController) Get() {
	this.activeContent("index", "")

	//******** This page requires login
	sess := this.GetSession("go.mbitson.com")
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
