package main

import (
	"fmt"
	"golang-CRUD/models"
	_ "golang-CRUD/routers"
	"os"
	"strconv"

	"github.com/astaxie/beego/client/orm"
	beego "github.com/astaxie/beego/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
}

func main() {
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.Debug = true
	o := orm.NewOrm()

	user := models.User{Username: "alailson", Email: "alailson@gmail.com", Password: "mypassword"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	dr := o.Driver()
	fmt.Println(dr.Name() == "default")    // true
	fmt.Println(dr.Type() == orm.DRSqlite) // true
	beego.BConfig.Listen.HTTPPort = getPort()
	beego.BConfig.RunMode = os.Getenv("ENVIRONMENT")
	beego.Run()
}

func getPort() int {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		portEnv, _ := strconv.Atoi("9090")
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
		return portEnv
	}
	portEnv, _ := strconv.Atoi(port)

	return portEnv
}
