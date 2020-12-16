package controllers

import (
	"fmt"
	"golang-CRUD/models"
	"log"

	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/core/validation"
	beego "github.com/astaxie/beego/server/web"
)

// SignInController logic for signup
type SignInController struct {
	beego.Controller
}

// Get the page
func (c *SignInController) Get() {
	// read flash message when making a request
	flash := beego.ReadFromRequest(&c.Controller)
	// logic for detect message of flash message
	if n, ok := flash.Data["error"]; ok {
		log.Println("some error", ok, n)
	}
	c.Data["content"] = "value"
	c.TplName = "login.tpl"
}

// Post receive data
func (c *SignInController) Post() {
	// define package
	o := orm.NewOrm()
	flash := beego.NewFlash()

	// get data from post method in form
	email := c.GetString("email")
	password := c.GetString("password")
	// organize data
	u := models.User{Email: email, Password: password}

	// interface for validation data
	valid := validation.Validation{}

	// validate email
	valid.Required(u.Email, "email")
	valid.Email(u.Email, "email")
	// validate password
	valid.Required(u.Password, "password")

	// verify if user exists
	var users []models.User

	var cond *orm.Condition
	cond = orm.NewCondition()
	cond = cond.And("email", email)
	var qs orm.QuerySeter

	qs = o.QueryTable("user").SetCond(cond)
	_, err := qs.All(&users)
	fmt.Println("password:", users[0].Password, "err:", err)

	// fmt.Println("queryTable email init")
	// fmt.Printf("Returned Rows Num: %v, %s\n", numPassword, errPassword)
	// fmt.Println("queryTable email end")
	// case user exist throws a error
	// if numEmail != 0 {
	// 	fmt.Println("email", numEmail)
	// 	valid.SetError("email", "already exists")
	// 	fmt.Println("email", numEmail)
	// }

	// show flash message case error
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			if err != nil {
				flash.Error("%s %s\n", err.Key, err.Message)
				flash.Store(&c.Controller)
				c.Redirect("/login", 302)
				return
			}
		}
	}
	// case all data is valid save in database
	// s := models.SaveUser(&u)

	// fmt.Println("saveuser is working:", s)
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	fmt.Println("user:", u)
	// redirect user to login
	c.Ctx.Redirect(200, "/login")
}
