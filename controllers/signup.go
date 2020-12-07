package controllers

import (
	"fmt"

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
	c.Data["content"] = "value"
	c.TplName = "register.tpl"
}

// Post receive data
func (c *SignUpController) Post() {
	username := c.GetString("username")
	email := c.GetString("email")
	password := c.GetString("password")
	passwordConfirm := c.GetString("passwordConfirm")
	fmt.Println("username:", username)
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	fmt.Println("passwordConfirm:", passwordConfirm)
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
