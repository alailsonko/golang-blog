package controllers

import (
	"fmt"
	"golang-CRUD/models"
	"log"

	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/core/validation"
	beego "github.com/astaxie/beego/server/web"
)

// SignUpController logic for signup
type SignUpController struct {
	beego.Controller
}

// Get the page
func (c *SignUpController) Get() {
	// read flash message when making a request
	flash := beego.ReadFromRequest(&c.Controller)
	// logic for detect message of flash message
	if n, ok := flash.Data["error"]; ok {
		log.Println("some error", ok, n)
	}
	c.Data["content"] = "value"
	c.TplName = "register.tpl"
}

// Post receive data
func (c *SignUpController) Post() {
	// define package
	o := orm.NewOrm()
	flash := beego.NewFlash()

	// get data from post method in form
	username := c.GetString("username")
	email := c.GetString("email")
	password := c.GetString("password")
	passwordConfirm := c.GetString("passwordConfirm")
	// organize data
	u := models.User{Username: username, Email: email, Password: password}

	// interface for validation data
	valid := validation.Validation{}

	// validate username
	valid.Required(u.Username, "username")
	// validate email
	valid.Required(u.Email, "email")
	valid.Email(u.Email, "email")
	// validate password
	valid.Required(u.Password, "password")
	valid.Required(passwordConfirm, "passwordConfirm")
	if u.Password != passwordConfirm {
		valid.SetError("password", "passwords does not match")
	}

	// verify if user exists
	var users []models.User
	numUsername, errUsername := o.QueryTable("user").Filter("username", u.Username).All(&users)
	numEmail, errEmail := o.QueryTable("user").Filter("email", u.Email).All(&users)
	fmt.Println("queryTable username init")
	fmt.Printf("Returned Rows Num: %v, %s\n", numUsername, errUsername)
	fmt.Println("queryTable username end")
	fmt.Println("queryTable email init")
	fmt.Printf("Returned Rows Num: %v, %s\n", numEmail, errEmail)
	fmt.Println("queryTable email end")

	// case user exist throws a error
	if numEmail != 0 {
		fmt.Println("email", numEmail, "username", numUsername)
		valid.SetError("email", "already exists")
		fmt.Println("email", numEmail, "username", numUsername)
	} else if numUsername != 0 {
		valid.SetError("username", "already exists")
	}

	// show flash message case error
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			if err != nil {
				flash.Error("%s %s\n", err.Key, err.Message)
				flash.Store(&c.Controller)
				c.Redirect("/register", 302)
				return
			}
		}
	}
	// case all data is valid save in database
	s := models.SaveUser(&u)

	fmt.Println("saveuser is working:", s)
	fmt.Println("username:", username)
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	fmt.Println("passwordConfirm:", passwordConfirm)
	fmt.Println("user:", u)
	// redirect user to login
	c.Ctx.Redirect(200, "/login")
}
