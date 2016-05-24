package controllers

import (
	"go.mbitson.com/models"
	"github.com/astaxie/beego/orm"
)

// operations for Sites.Go
type EventsApiController struct {
	MainController
}

func (this *EventsApiController) URLMapping() {
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
func (this *EventsApiController) Post() {

}

// @Title Get
// @Description get Sites.Go by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is empty
// @router /:id [get]
func (this *EventsApiController) GetOne() {

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
func (this *EventsApiController) GetAll() {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load ORM
	o := orm.NewOrm()
	o.Using("default")

	var events []*models.Log
	o.QueryTable("log").Filter("User_id", m["id"]).OrderBy("-Created").Limit(10).All(&events)
	this.ajaxContent(&events)
}

func (this *EventsApiController) AddNew() bool {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return false
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Load orm
	o := orm.NewOrm()
	o.Using("default")

	site := models.Site{Domain: "go.mbitson.com", User: &models.AuthUser{Id:m["id"].(int)}}

	_, err := o.Insert(&site)
	if err != nil {
		return true
	}

	return false
}

// @Title Update
// @Description update the Sites.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sites.Go	true		"body for Sites.Go content"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *EventsApiController) Put() {
	
}

// @Title Delete
// @Description delete the Sites.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *EventsApiController) Delete() {
	
}
