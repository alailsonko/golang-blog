package main

import (
	_ "golang-CRUD/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

