package controllers

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
	"time"
)

// operations for Alerts.Go
type AlertsApiController struct {
	MainController
}

func (this *AlertsApiController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Alerts.Go
// @Param	body		body 	models.Alerts.Go	true		"body for Alerts.Go content"
// @Success 200 {int} models.Alerts.Go.Id
// @Failure 403 body is empty
// @router / [post]
func (this *AlertsApiController) Post() {
	//	// Load session data
	//	sess := this.GetSession("os-auth")
	//	if sess == nil {
	//		this.Redirect("/user/login/", 302)
	//		return
	//	}
	//
	//	// Map userdata to m
	//	m := sess.(map[string]interface{})
	//
	//	// Load orm
	//	o := orm.NewOrm()
	//	o.Using("default")
	//
	//	// Get the post data
	//	s_id := this.Ctx.Input.Params[":siteId"]
	//	siteId, _ := strconv.ParseInt(s_id, 0, 64)
	//	interval_str := this.Ctx.Input.Params[":interval"]
	//	interval, _ := strconv.ParseInt(interval_str, 0, 64)
	//	t_id := this.Ctx.Input.Params[":typeId"]
	//	typeId, _ := strconv.ParseInt(t_id, 0, 64)
	//	alertType := models.AlertType{Id: int(typeId)}
	//
	//	// Get this site from DB to confirm access for this user.
	//	var site *models.Site
	//	err := o.QueryTable("site").Filter("Id", siteId).Filter("User_id", m["id"]).One(&site)
	//	if err == orm.ErrNoRows {
	//		this.ajaxContent("You do not have access to create a alert for this site.")
	//		return
	//	}
	//
	//	// Attempt insert
	//	alert := models.Alert{Site: site, Interval: int(interval), Alert_type: &alertType}
	//	m_id, err := o.Insert(&alert)
	//	alert_id := int(m_id)
	//	if err != nil {
	//		this.ajaxContent("Error adding alert to system.")
	//		return
	//	}
	//
	//	// Log this creation
	//	this.logCustomActionAndId("Alert.Create", alert_id, m["id"].(int))
	//
	//	// If no error
	//	if err == nil {
	//		this.Data["json"] = alert_id
	//	}
	//	this.ServeJSON()
}

func (this *AlertsApiController) Get() {
	alertId := this.Ctx.Input.Param(":objectId")
	if alertId != "" {
		this.GetOne()
	} else {
		this.GetAll()
	}
}

// @Title Get
// @Description get Alerts.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Alerts.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *AlertsApiController) GetOne() {
	// Get Alert ID
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

	var alerts []*models.Alert
	_, err := o.QueryTable("alert").RelatedSel().Filter("Site__Id", siteId).Filter("Site__User__Id", m["id"]).All(&alerts)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	} else {
		this.ajaxContent(&alerts)
	}
}

func (this *AlertsApiController) GetRecent() {
	// Get Alert ID
	siteId := this.GetString("id")

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

	var alerts []*models.Alert
	_, err := o.
		QueryTable("alert").
		RelatedSel().
		Filter("Site__Id", siteId).
		Filter("Site__User", &models.AuthUser{Id: m["id"].(int)}).
		Filter("Checked_date__gte", time.Now().Add(-24*time.Hour)).
		All(&alerts)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	} else {
		this.ajaxContent(&alerts)
	}
}

// @Title Get All
// @Description get Alerts.Go
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Alerts.Go
// @Failure 403
// @router / [get]
func (this *AlertsApiController) GetAll() {
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

	var alerts []*models.Alert
	o.QueryTable("alert").Filter("User_id", m["id"]).All(&alerts)
	this.ajaxContent(&alerts)
}

// @Title Update
// @Description update the Alerts.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Alerts.Go	true		"body for Alerts.Go content"
// @Success 200 {object} models.Alerts.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *AlertsApiController) Put() {

}

// @Title Delete
// @Description delete the Alerts.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *AlertsApiController) Delete() {
	// Get Alert ID
	alertId := this.Ctx.Input.Param(":objectId")

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	// Load Alert
	var alert models.Alert
	alert_err := o.QueryTable("alert").Filter("Id", alertId).One(&alert)
	if alert_err == orm.ErrNoRows {
		this.ajaxContent("No Alert Found With This ID")
		return
	}

	// Determine access to site
	access := this.canAccessSite(alert.Site.Id)
	if access == false {
		this.ajaxContent("This user does not have access to delete this alert.")
		return
	}

	// Delete alert
	_, delete_err := o.QueryTable("alert").Filter("Id", alertId).Delete()
	if delete_err != nil {
		this.ajaxContent("Could Not Delete Alert!")
		return
	}

	// Return success!
	this.ajaxContent("Success")
}

func (this *AlertsApiController) canAccessSite(site_id int) bool {
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

func (this *AlertsApiController) getUserData() (map[string]interface{}, error) {
	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return nil, errors.New("Not Logged In")
	}

	return sess.(map[string]interface{}), nil
}
