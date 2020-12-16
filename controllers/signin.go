package controllers

import (
	"fmt"
	"golang-CRUD/models"
	"log"

	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/core/validation"
	beego "github.com/astaxie/beego/server/web"
	"golang.org/x/crypto/bcrypt"
)

// SignInController logic for signup
type SignInController struct {
	beego.Controller
}

// Get the page
func (c *SignInController) Get() {
	// read flash message when making a request
	v := c.GetSession("sonko")
	if v != nil {
		c.Redirect("/", 200)
		return
	}
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
	// fmt.Println(qs)
	us, err := qs.All(&users)
	fmt.Println(us)
	fmt.Println(err)
	if us == 1 {
		fmt.Println("password:", users[0].Password, "err:", err)
		err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(u.Password))
		fmt.Println("err: compare:", err)
		if err != nil {
			fmt.Println("user", users)
			valid.SetError("user:", "credentials error")
			fmt.Println("user", users)
		}
		if err == nil {
			c.SetSession("sonko", int(1))
			c.Data["num"] = 0
		}
	}

	if us == 0 {
		fmt.Println("user", users)
		valid.SetError("user:", "credentials error")
		fmt.Println("user", users)
	}

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
	c.Ctx.Redirect(200, "/")
}
