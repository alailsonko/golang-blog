package controllers

import (
	"fmt"

	beego "github.com/astaxie/beego/server/web"
)

// MainController type
type MainController struct {
	beego.Controller
}

// Get the page
func (c *MainController) Get() {

	v := c.GetSession("sonko")
	fmt.Println("getsession:", v)
	if v == nil {
		fmt.Println("getsession:", v)

		c.Ctx.Redirect(401, "/login")
		return
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// Register page
func (c *MainController) Register() {

	c.Data["Website"] = "alailson"
	c.Data["Email"] = "aalailson3@gmail.com"
	c.TplName = "register.tpl"
}
