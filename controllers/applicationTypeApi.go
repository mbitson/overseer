package controllers

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
	//	"strconv"
)

// operations for Monitors.Go
type ApplicationTypesApiController struct {
	MainController
}

func (this *ApplicationTypesApiController) URLMapping() {
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
func (this *ApplicationTypesApiController) Post() {
	// Cannot add one.
}

func (this *ApplicationTypesApiController) Get() {
	id := this.Ctx.Input.Param(":objectId")
	if id != "" {
		this.GetOne()
	}
	this.GetAll()
}

// @Title Get
// @Description get Monitors.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Monitors.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *ApplicationTypesApiController) GetOne() {
	//	// Get Monitor ID
	//	siteId := this.GetString("id")
	//	start := this.GetString("start")
	//	end := this.GetString("end")
	//
	//	// Load session data
	//	sess := this.GetSession("os-auth")
	//	if sess == nil {
	//		this.Redirect("/user/login/", 302)
	//		return
	//	}
	//
	//	// Map userdata to m
	//	m := sess.(map[string]interface{})
	//	user_id := strconv.Itoa(m["id"].(int))
	//
	//	// Load ORM
	//	o := orm.NewOrm()
	//	o.Using("default")
	//
	//	// Build query params
	//	queryParams := []string{siteId, user_id, start, end}
	//	var runs []orm.Params
	//	o.Raw(
	//		"SELECT FORMAT(AVG(runs.ping), 0) as Ping, FORMAT(AVG(runs.response_time), 0) as Response_time, MAX(runs.status_code) as Status_code, FROM_UNIXTIME( TRUNCATE(UNIX_TIMESTAMP(runs.time_run) / 600,0)*600) AS Time_run FROM monitor_run AS runs JOIN monitor as monitors ON monitors.id = runs.monitor_id JOIN site as sites ON sites.id = monitors.site_id WHERE sites.id = ? AND sites.user_id = ? AND runs.time_run >= ? AND runs.time_run < ? GROUP BY FROM_UNIXTIME( TRUNCATE(UNIX_TIMESTAMP(runs.time_run) / 600,0)*600)",
	//		queryParams,
	//	).Values(&runs)

	//	qb.Select("Ping", "Response_time", "Status_code", "Time_run", "Status_code", "ROUND(UNIX_TIMESTAMP(timestamp)/(15 * 60)) AS timekey").
	//		From("monitor_run").
	//		Where("Monitor__Site__Id", siteId).
	//		Where("Monitor__Site__User_id", m["id"]).
	//		Where("Time_run__gte", start).
	//		Where("Time_run__lt", end).
	//		All(&runs, )

	//	var runs []*models.ApplicationType
	//	o.
	//		QueryTable("monitor_run").
	//		Filter("Monitor__Site__Id", siteId).
	//		Filter("Monitor__Site__User_id", m["id"]).
	//		Filter("Time_run__gte", start).
	//		Filter("Time_run__lt", end).
	//		All(&runs, "Ping", "Response_time", "Status_code", "Time_run", "Status_code")
	//	this.ajaxContent(&runs)
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
func (this *ApplicationTypesApiController) GetAll() {
	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var types []*models.ApplicationType
	o.QueryTable("application_type").All(&types)
	this.ajaxContent(&types)
}

// @Title Update
// @Description update the Monitors.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Monitors.Go	true		"body for Monitors.Go content"
// @Success 200 {object} models.Monitors.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *ApplicationTypesApiController) Put() {

}

// @Title Delete
// @Description delete the Monitors.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ApplicationTypesApiController) Delete() {
	//	// Get Monitor ID
	//	monitorId := this.Ctx.Input.Param(":objectId")
	//
	//	// Load ORM
	//	o := orm.NewOrm()
	//	o.Using("default")
	//
	//	// Load Monitor
	//	var monitor models.Monitor
	//	monitor_err := o.QueryTable("monitor").Filter("Id", monitorId).One(&monitor)
	//	if monitor_err == orm.ErrNoRows {
	//		this.ajaxContent("No Monitor Found With This ID")
	//		return
	//	}
	//
	//	// Determine access to site
	//	access := this.canAccessSite(monitor.Site.Id)
	//	if access == false {
	//		this.ajaxContent("This user does not have access to delete this monitor.")
	//		return
	//	}
	//
	//	// Delete monitor
	//	_, delete_err := o.QueryTable("monitor").Filter("Id", monitorId).Delete()
	//	if delete_err != nil{
	//		this.ajaxContent("Could Not Delete Monitor!")
	//		return
	//	}
	//
	//	// Return success!
	//	this.ajaxContent("Success")
}

func (this *ApplicationTypesApiController) getUserData() (map[string]interface{}, error) {
	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return nil, errors.New("Not Logged In")
	}

	return sess.(map[string]interface{}), nil
}
