package controllers

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
)

// operations for Monitors.Go
type MonitorRunsApiController struct {
	MainController
}

func (this *MonitorRunsApiController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Monitors.Go
// @Param	body		body 	models.Monitors.Go	true		"body for Monitors.Go content"
// @Success 200 {int} models.Monitors.Go.Id
// @Failure 403 body is empty
// @router / [post]
func (this *MonitorRunsApiController) Post() {
	// Cannot add one.
}

func (this *MonitorRunsApiController) Get() {
	siteId := this.Ctx.Input.Param(":objectId")
	if siteId != "" {
		this.GetOne()
	}
}

// @Title Get
// @Description get Monitors.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Monitors.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *MonitorRunsApiController) GetOne() {
	// Get Monitor ID
	siteId := this.Ctx.Input.Param(":id")
	start := this.Ctx.Input.Param(":start")
	end := this.Ctx.Input.Param(":end")

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

	// Build query params
	var runs []*models.MonitorRun
	o.
		QueryTable("monitor_run").
		Filter("Monitor__Site__User__Id", m["id"]).
		Filter("Monitor__Site__Id", siteId).
		Filter("Time_run__gte", start).
		Filter("Time_run__lt", end).
		All(&runs, "Ping", "Response_time", "Status_code", "Time_run", "Status_code")
	this.ajaxContent(&runs)
}

// @Title Get All
// @Description get Monitors.Go
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Monitors.Go
// @Failure 403
// @router / [get]
func (this *MonitorRunsApiController) GetAll() {
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

	var monitors []*models.Monitor
	o.QueryTable("monitor").Filter("User_id", m["id"]).All(&monitors)
	this.ajaxContent(&monitors)
}

// @Title Update
// @Description update the Monitors.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Monitors.Go	true		"body for Monitors.Go content"
// @Success 200 {object} models.Monitors.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *MonitorRunsApiController) Put() {

}

// @Title Delete
// @Description delete the Monitors.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *MonitorRunsApiController) Delete() {
	// Get Monitor ID
	monitorId := this.Ctx.Input.Param(":objectId")

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	// Load Monitor
	var monitor models.Monitor
	monitor_err := o.QueryTable("monitor").Filter("Id", monitorId).One(&monitor)
	if monitor_err == orm.ErrNoRows {
		this.ajaxContent("No Monitor Found With This ID")
		return
	}

	// Determine access to site
	access := this.canAccessSite(monitor.Site.Id)
	if access == false {
		this.ajaxContent("This user does not have access to delete this monitor.")
		return
	}

	// Delete monitor
	_, delete_err := o.QueryTable("monitor").Filter("Id", monitorId).Delete()
	if delete_err != nil {
		this.ajaxContent("Could Not Delete Monitor!")
		return
	}

	// Return success!
	this.ajaxContent("Success")
}

func (this *MonitorRunsApiController) canAccessSite(site_id int) bool {
	// Map userdata to m
	m, err := this.getUserData()
	if err != nil {
		return false
	}

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	// Load site
	var site models.Site
	site_err := o.QueryTable("site").Filter("Id", site_id).Filter("User_id", m["id"]).One(&site)
	if site_err == orm.ErrNoRows {
		return false
	}

	return true
}

func (this *MonitorRunsApiController) getUserData() (map[string]interface{}, error) {
	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return nil, errors.New("Not Logged In")
	}

	return sess.(map[string]interface{}), nil
}
