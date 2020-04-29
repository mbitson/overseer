package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/mbitson/overseer/models"
)

// operations for Sites.Go
type ContactsApiController struct {
	MainController
}

func (this *ContactsApiController) URLMapping() {
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
func (this *ContactsApiController) Post() {

	// Load session data
	sess := this.GetSession("os-auth")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Get posted data
	first := this.GetString("First")
	last := this.GetString("Last")
	email := this.GetString("Email")
	sms := this.GetString("Sms")

	// Load orm
	o := orm.NewOrm()
	o.Using("default")

	// Insert contact!
	contact := models.Contact{User: &models.AuthUser{Id: m["id"].(int)}, First: first, Last: last, Email: email, Sms: sms}
	contact_id, contactInsertError := o.Insert(&contact)
	this.logCustomActionAndId("Contact.Create", int(contact_id), m["id"].(int))

	// If no error
	if contactInsertError == nil {
		this.Data["json"] = contact
	} else {
		this.Data["json"] = "Error inserting contact."
	}
	this.ServeJSON()
}

func (this *ContactsApiController) Get() {
	contactId := this.Ctx.Input.Param(":objectId")
	if contactId != "" {
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
func (this *ContactsApiController) GetOne() {
	// Get Contact ID
	contactId := this.Ctx.Input.Param(":objectId")

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

	var contacts models.Contact
	err := o.QueryTable("contact").Filter("Id", contactId).Filter("User", &models.AuthUser{Id: m["id"].(int)}).One(&contacts)
	if err == orm.ErrNoRows {
		this.ajaxContent("Failure")
	} else {
		this.ajaxContent(&contacts)
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
func (this *ContactsApiController) GetAll() {
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

	var contacts []orm.Params
	o.QueryTable("contact").Filter("User", &models.AuthUser{Id: m["id"].(int)}).Values(&contacts)
	this.ajaxContent(&contacts)
}

// @Title Update
// @Description update the Sites.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sites.Go	true		"body for Sites.Go content"
// @Success 200 {object} models.Sites.Go
// @Failure 403 :id is not int
// @router /:id [put]
func (this *ContactsApiController) Put() {

}

// @Title Delete
// @Description delete the Sites.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *ContactsApiController) Delete() {
	// Get Contact ID
	contactId := this.Ctx.Input.Param(":objectId")

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

	o.QueryTable("contact").Filter("User", &models.AuthUser{Id: m["id"].(int)}).Filter("Id", contactId).Delete()
	this.ajaxContent("Success")
}
