package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
)

// operations for Sites.Go
type SitesApiController struct {
	MainController
}

func (this *SitesApiController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Sites.Go
// @Param	body		body 	models.Sites.Go	true		"body for Sites.Go content"
// @Success 200 {int} models.Sites.Go.Id
// @Failure 403 body is empty
// @router / [post]
func (this *SitesApiController) Post() {

	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Get posted data
	domain := this.GetString("Domain")

	// Load orm
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	// Attempt insert
	site := models.Site{Domain: domain, User: &models.AuthUser{Id: m["id"].(int)}}
	id, siteInsertError := o.Insert(&site)
	this.logCustomActionAndId("Site.Create", int(id), m["id"].(int))

	// Add monitor and log.
	monitorType := models.MonitorType{Id: int(1)}
	monitor := models.Monitor{Name: "Default", Site: &site, Interval: 60, Monitor_type: &monitorType}
	m_id, monitorInsertError := o.Insert(&monitor)
	this.logCustomActionAndId("Monitor.Create", int(m_id), m["id"].(int))

	// Add default alert and log.
	alert := models.Alert{Name: "Default Alert", Site: &site, Email: "me@mbitson.com", Sms: "2313296944", Urgency: 2}
	a_id, alertInsertError := o.Insert(&alert)
	this.logCustomActionAndId("Alert.Create", int(a_id), m["id"].(int))

	// If no error
	if siteInsertError == nil && monitorInsertError == nil && alertInsertError == nil {
		// Commit changes!
		o.Commit()
		// Override failure response with success message!
		this.Data["json"] = site
	} else {
		o.Rollback()
		this.Data["json"] = "Error adding your site. Please contact support."
	}
	this.ServeJSON()
}

func (this *SitesApiController) Get() {
	siteId := this.Ctx.Input.Param(":objectId")
	if siteId != "" {
		this.GetOne()
	} else {
		this.GetAll()
	}
}

// @Title Get
// @Description get Sites.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *SitesApiController) GetOne() {
	// Get Site ID
	siteId := this.Ctx.Input.Param(":objectId")

	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var sites models.Site
	err := o.QueryTable("site").Filter("Id", siteId).Filter("User", &models.AuthUser{Id: m["id"].(int)}).One(&sites)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	} else {
		this.ajaxContent(&sites)
	}
}

// @Title Get All
// @Description get Sites.Go
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Sites.Go
// @Failure 403
// @router / [get]
func (this *SitesApiController) GetAll() {
	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var sites []orm.Params
	o.QueryTable("site").Filter("User", &models.AuthUser{Id: m["id"].(int)}).Values(&sites)
	//o.Raw("SELECT site.id, site.domain, site.reg_date, site.user_id FROM site AS site JOIN monitor AS monitor ON site.id == monitor.id JOIN monitor_run AS run ON run.monitor_id == monitor.id WHERE site.user_id = ?", m["id"]).Values(&sites)
	this.ajaxContent(&sites)
}

// @Title Update
// @Description update the Sites.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sites.Go	true		"body for Sites.Go content"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *SitesApiController) Put() {

}

// @Title Delete
// @Description delete the Sites.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *SitesApiController) Delete() {
	// Get Site ID
	siteId := this.Ctx.Input.Param(":objectId")

	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	o.QueryTable("site").Filter("User", &models.AuthUser{Id: m["id"].(int)}).Filter("Id", siteId).Delete()
	this.ajaxContent("Success")
}
