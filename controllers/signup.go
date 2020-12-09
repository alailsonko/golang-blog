package controllers

import (
	"fmt"
	"golang-CRUD/models"
	"log"

	"github.com/astaxie/beego/core/validation"
	beego "github.com/astaxie/beego/server/web"
)

// SignUpController logic for signup
type SignUpController struct {
	beego.Controller
}

//Prepare for add to database
func (c *SignUpController) Prepare() {

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

	// TODO validate the data
	// TODO handle errors
	// TODO a model
	// TODO save data in Database
	// TODO cryptograph the data
	flash := beego.NewFlash()

	username := c.GetString("username")
	email := c.GetString("email")
	password := c.GetString("password")
	passwordConfirm := c.GetString("passwordConfirm")
	u := models.User{Username: username, Email: email, Password: password}
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
	s := models.SaveUser(&u)
	fmt.Println("saveuser is working:", s)
	fmt.Println("username:", username)
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	fmt.Println("passwordConfirm:", passwordConfirm)
	fmt.Println("user:", u)
	// pk := models.GetCruPkg(pkgname)
	// if pk.Id == 0 {
	//     var pp models.PkgEntity
	//     pp.Pid = 0
	//     pp.Pathname = pkgname
	//     pp.Intro = pkgname
	//     models.InsertPkg(pp)
	//     pk = models.GetCruPkg(pkgname)
	// }
	// var at models.Article
	// at.Pkgid = pk.Id
	// at.Content = content
	// models.InsertArticle(at)
	c.Ctx.Redirect(302, "/")
}
