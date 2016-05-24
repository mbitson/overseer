package controllers

import (
	"go.mbitson.com/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"errors"
)

// operations for Monitors.Go
type MonitorsApiController struct {
	MainController
}

func (this *MonitorsApiController) URLMapping() {
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
func (this *MonitorsApiController) Post() {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load orm
	o := orm.NewOrm()
	o.Using("default")

	// Get the post data
	s_id := this.Ctx.Input.Params[":siteId"]
	siteId, _ := strconv.ParseInt(s_id, 0, 64)
	interval_str := this.Ctx.Input.Params[":interval"]
	interval, _ := strconv.ParseInt(interval_str, 0, 64)
	t_id := this.Ctx.Input.Params[":typeId"]
	typeId, _ := strconv.ParseInt(t_id, 0, 64)
	monitorType := models.MonitorType{Id: int(typeId)}

	// Get this site from DB to confirm access for this user.
	var site *models.Site
	err := o.QueryTable("site").Filter("Id", siteId).Filter("User_id", m["id"]).One(&site)
	if err == orm.ErrNoRows {
		this.ajaxContent("You do not have access to create a monitor for this site.")
		return
	}

	// Attempt insert
	monitor := models.Monitor{Site: site, Interval: int(interval), Monitor_type: &monitorType}
	m_id, err := o.Insert(&monitor)
	monitor_id := int(m_id)
	if err != nil {
		this.ajaxContent("Error adding monitor to system.")
		return
	}

	// Log this creation
	this.logCustomActionAndId("Monitor.Create", monitor_id, m["id"].(int))

	// If no error
	if err == nil {
		this.Data["json"] = monitor_id
	}
	this.ServeJson()
}

func (this *MonitorsApiController) Get() {
	monitorId := this.Ctx.Input.Params[":objectId"]
	if monitorId != "" {
		this.GetOne()
	} else {
		this.GetAll()
	}
}

// @Title Get
// @Description get Monitors.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Monitors.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *MonitorsApiController) GetOne() {
	// Get Monitor ID
	siteId := this.Ctx.Input.Params[":objectId"]

	// Load session data
	sess := this.GetSession("go.mbitson.com")
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
	_, err := o.QueryTable("monitor").RelatedSel().Filter("Site__Id", siteId).Filter("Site__User__Id", m["id"]).All(&monitors)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	}else{
		this.ajaxContent(&monitors)
	}
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
func (this *MonitorsApiController) GetAll() {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
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
func (this *MonitorsApiController) Put() {
	
}

// @Title Delete
// @Description delete the Monitors.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *MonitorsApiController) Delete() {
	// Get Monitor ID
	monitorId := this.Ctx.Input.Params[":objectId"]

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
	if delete_err != nil{
		this.ajaxContent("Could Not Delete Monitor!")
		return
	}

	// Return success!
	this.ajaxContent("Success")
}

func (this *MonitorsApiController) canAccessSite(site_id int) (bool) {
	// Map userdata to m
	m, err := this.getUserData()
	if err != nil{
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

func (this *MonitorsApiController) getUserData() (map[string]interface{}, error) {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return nil, errors.New("Not Logged In")
	}

	return sess.(map[string]interface{}), nil
}
