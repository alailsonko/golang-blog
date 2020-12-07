package controllers

import (
	beego "github.com/astaxie/beego/server/web"
)

// MainController type
type MainController struct {
	beego.Controller
}

// Get the page
func (c *MainController) Get() {
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
