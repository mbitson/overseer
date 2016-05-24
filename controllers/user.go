package controllers

import (
	"go.mbitson.com/models"
	pk "go.mbitson.com/utilities/pbkdf2"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v1"
	"crypto/tls"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/twinj/uuid"
	"strings"
	"time"
)

type UserController struct {
	MainController
}


func (this *UserController) Login() {
	this.bareContent("user/login", "")

	back := strings.Replace(this.Ctx.Input.Param(":back"), ">", "/", -1) // allow for deeper URL such as l1/l2/l3 represented by l1>l2>l3
	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		email := this.GetString("email")
		password := this.GetString("password")
		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.Required(password, "password")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		//******** Read password hash from database
		var x pk.PasswordHash

		x.Hash = make([]byte, 32)
		x.Salt = make([]byte, 16)

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: email}
		err := o.Read(&user, "Email")
		if err == nil {
			if user.Reg_key != "" {
				flash.Error("Account not verified")
				flash.Store(&this.Controller)
				return
			}

			// scan in the password hash/salt
			if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
				fmt.Println("ERROR:", err)
			}
			if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
				fmt.Println("ERROR:", err)
			}
		} else {
			flash.Error("No such user/email")
			flash.Store(&this.Controller)
			return
		}

		//******** Compare submitted password with database
		if !pk.MatchPassword(password, &x) {
			flash.Error("Bad password")
			flash.Store(&this.Controller)
			return
		}

		//******** Create session and go back to previous page
		m := make(map[string]interface{})
		m["first"] = user.First
		m["id"] = user.Id
		m["username"] = email
		m["timestamp"] = time.Now()

		// Save user into session
		this.SetSession("go.mbitson.com", m)

		// Log this login
		this.logRouteAction(m["id"].(int))

		// Go back to their page!
		this.Redirect("/"+back, 302)
	}
}

func (this *UserController) Logout() {
	// Load session data
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}

	// Map userdata to m
	m := sess.(map[string]interface{})

	// Set view
	this.activeContent("logout", "")

	// Log this logout
	this.logRouteAction(m["id"].(int))

	// Delete session data for logged in user
	this.DelSession("go.mbitson.com")

	// Redirect to homepage
	this.Redirect("/", 302)
}

func (this *UserController) Register() {
	this.bareContent("user/register", "")

	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		first := this.GetString("first_name")
		last := this.GetString("last_name")
		email := this.GetString("email")
		password := this.GetString("password")
		password2 := this.GetString("password_2")
		valid := validation.Validation{}
		valid.Required(first, "first")
		valid.Email(email, "email")
		valid.MinSize(password, 6, "password")
		valid.Required(password2, "password_2")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}
		if password != password2 {
			flash.Error("Passwords don't match")
			flash.Store(&this.Controller)
			return
		}
		h := pk.HashPassword(password)

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		user := models.AuthUser{First: first, Last: last, Email: email}

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)

		// Add user to database with new uuid and send verification email
		u := uuid.NewV4()
		user.Reg_key = u.String()
		u_id, err := o.Insert(&user)
		if err != nil {
			flash.Error(email + " already registered")
			flash.Store(&this.Controller)
			return
		}

		// Cast user id.
		var user_id int
		user_id = int(u_id)

		// Log this logout
		this.logRouteAction(user_id)

		if !sendVerification(email, u.String()) {
			flash.Error("Unable to send verification email")
			flash.Store(&this.Controller)
			return
		}
		flash.Notice("Your account has been created. You must verify the account in your email.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}

func sendVerification(email, u string) bool{
	// Create a new message, configure headers.
	msg := gomail.NewMessage()
	msg.SetHeader("From", "mikelbitson@gmail.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Account Verification for go.mbitson.com")

	// Build link and body
	link := "http://go.mbitson.com/user/verify/" + u
	msg.SetBody("text/html", "To verify your account, please click on the link: <a href=\""+link+
				"\">"+link+"</a><br><br>Best Regards,<br>Mikel Bitson")

	// Create mailer for our message
	m := gomail.NewMailer(
		"mbitson.com",
		"go@mbitson.com",
		"~paran01d~",
		465, gomail.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)

	// Attempt send
	if err := m.Send(msg); err != nil {
		return false
	}
	return true
}

func (this *UserController) Verify() {
	this.activeContent("user/verify", "")

	u := this.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Reg_key: u}
	err := o.Read(&user, "Reg_key")
	if err == nil {
		this.Data["Verified"] = 1
		user.Reg_key = ""
		if _, err := o.Update(&user); err != nil {
			delete(this.Data, "Verified")
		}
	}
}

func (this *UserController) Profile() {
	this.activeContent("user/profile", "")

	//******** This page requires login
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}
	m := sess.(map[string]interface{})

	flash := beego.NewFlash()

	//******** Read password hash from database
	var x pk.PasswordHash

	x.Hash = make([]byte, 32)
	x.Salt = make([]byte, 16)

	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Email: m["username"].(string)}
	err := o.Read(&user, "Email")
	if err == nil {
		// scan in the password hash/salt
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}
	} else {
		flash.Error("Internal error")
		flash.Store(&this.Controller)
		return
	}

	// this deferred function ensures that the correct fields from the database are displayed
	defer func(this *UserController, user *models.AuthUser) {
		this.Data["First"] = user.First
		this.Data["Last"] = user.Last
		this.Data["Email"] = user.Email
	}(this, &user)

	if this.Ctx.Input.Method() == "POST" {
		first := this.GetString("first")
		last := this.GetString("last")
		email := this.GetString("email")
		current := this.GetString("current")
		password := this.GetString("password")
		password2 := this.GetString("password2")
		valid := validation.Validation{}
		valid.Required(first, "first")
		valid.Email(email, "email")
		valid.Required(current, "current")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		if password != "" {
			valid.MinSize(password, 6, "password")
			valid.Required(password2, "password2")
			if valid.HasErrors() {
				errormap := []string{}
				for _, err := range valid.Errors {
					errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
				}
				this.Data["Errors"] = errormap
				return
			}

			if password != password2 {
				flash.Error("Passwords don't match")
				flash.Store(&this.Controller)
				return
			}
			h := pk.HashPassword(password)

			// Convert password hash to string
			user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
		}

		//******** Compare submitted password with database
		if !pk.MatchPassword(current, &x) {
			flash.Error("Bad current password")
			flash.Store(&this.Controller)
			return
		}

		//******** Save user info to database
		user.First = first
		user.Last = last
		user.Email = email

		_, err := o.Update(&user)
		if err == nil {
			flash.Notice("Profile updated")
			flash.Store(&this.Controller)
			m["username"] = email
		} else {
			flash.Error("Internal error")
			flash.Store(&this.Controller)
			return
		}
	}
}

func (this *UserController) Remove() {
	this.activeContent("user/remove", "")

	//******** This page requires login
	sess := this.GetSession("go.mbitson.com")
	if sess == nil {
		this.Redirect("/user/login/", 302)
		return
	}
	m := sess.(map[string]interface{})

	if this.Ctx.Input.Method() == "POST" {
		current := this.GetString("current")
		valid := validation.Validation{}
		valid.Required(current, "current")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		flash := beego.NewFlash()

		//******** Read password hash from database
		var x pk.PasswordHash

		x.Hash = make([]byte, 32)
		x.Salt = make([]byte, 16)

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: m["username"].(string)}
		err := o.Read(&user, "Email")
		if err == nil {
			// scan in the password hash/salt
			if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
				fmt.Println("ERROR:", err)
			}
			if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
				fmt.Println("ERROR:", err)
			}
		} else {
			flash.Error("Internal error")
			flash.Store(&this.Controller)
			return
		}

		//******** Compare submitted password with database
		if !pk.MatchPassword(current, &x) {
			flash.Error("Bad current password")
			flash.Store(&this.Controller)
			return
		}

		//******** Delete user record
		_, err = o.Delete(&user)
		if err == nil {
			flash.Notice("Your account is deleted.")
			flash.Store(&this.Controller)
			this.DelSession("go.mbitson.com")
			this.Redirect("/notice", 302)
		} else {
			flash.Error("Internal error")
			flash.Store(&this.Controller)
			return
		}
	}
}
