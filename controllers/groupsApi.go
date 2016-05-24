package controllers

import (
	"go.mbitson.com/models"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"encoding/json"
)

// operations for Sites.Go
type GroupsApiController struct {
	MainController
}

func (this *GroupsApiController) URLMapping() {
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
func (this *GroupsApiController) Post() {

	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Get posted data
	name := this.GetString("Name")
	contacts := this.Input().Get("Contacts")

	// Load orm
	o := orm.NewOrm()
	o.Using("default")

	// Insert group!
	group := models.Group{User: &models.AuthUser{Id:m["id"].(int)}, Name: name, Contacts: contacts}
	group_id, groupInsertError := o.Insert(&group)
	this.logCustomActionAndId("Group.Create", int(group_id), m["id"].(int))

	// If no error
	if groupInsertError == nil {
		this.Data["json"] = group
	}else{
		this.Data["json"] = "Error inserting group."
	}
	this.ServeJson()
}
func (this *GroupsApiController) SaveContacts() {

//	// Load session data
//	sess := this.GetSession("go.mbitson.com")
//	if sess == nil {
//		this.Redirect("/user/login/", 302)
//		return
//	}
//
//	// Map userdata to m
//	m := sess.(map[string]interface{})

	var groups []models.Group
	json.Unmarshal(this.Ctx.Input.RequestBody, &groups)
	spew.Dump(this.Ctx.Input.RequestBody)


	// Load orm
	o := orm.NewOrm()
	o.Using("default")

	for _, group := range groups{
		o.Update(&group)
	}

	// Insert group!
//	group := models.Group{User: &models.AuthUser{Id:m["id"].(int)}, Name: name, Contacts: contacts}
//	group_id, groupInsertError := o.Insert(&group)
//	this.logCustomActionAndId("Group.Create", int(group_id), m["id"].(int))

	// If no error
	//if groupInsertError == nil {
		this.Data["json"] = groups
//	}else {
//		this.Data["json"] = "Error Saving Groups."
//	}
	this.ServeJson()
}

func (this *GroupsApiController) Get() {
	groupId := this.Ctx.Input.Params[":objectId"]
	if groupId != "" {
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
func (this *GroupsApiController) GetOne() {
	// Get Group ID
	groupId := this.Ctx.Input.Params[":objectId"]

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

	var groups models.Group
	err := o.QueryTable("group").Filter("Id", groupId).Filter("User", &models.AuthUser{Id:m["id"].(int)}).One(&groups)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	}else{
		this.ajaxContent(&groups)
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
func (this *GroupsApiController) GetAll() {
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

	var groups []models.Group
	o.QueryTable("group").Filter("User", &models.AuthUser{Id:m["id"].(int)}).All(&groups)
	for i := 0; i < len(groups); i++ {
		var contacts []models.Contact
		o.Raw(fmt.Sprint("SELECT id, first, last, email, sms FROM contact WHERE id IN(", groups[i].Contacts, ")")).QueryRows(&contacts)
		groups[i].ContactDetails = &contacts
	}
	this.ajaxContent(&groups)
}

// @Title Update
// @Description update the Sites.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sites.Go	true		"body for Sites.Go content"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *GroupsApiController) Put() {
	
}

// @Title Delete
// @Description delete the Sites.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *GroupsApiController) Delete() {
	// Get Group ID
	groupId := this.Ctx.Input.Params[":objectId"]

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

	o.QueryTable("group").Filter("User", &models.AuthUser{Id:m["id"].(int)}).Filter("Id", groupId).Delete()
	this.ajaxContent("Success")
}
