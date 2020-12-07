package routers

import (
	"golang-CRUD/src/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
