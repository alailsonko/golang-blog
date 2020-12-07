package main

import (
	"fmt"
	_ "golang-CRUD/routers"
	"os"
	"strconv"

	beego "github.com/astaxie/beego/server/web"
)

func main() {

	beego.BConfig.Listen.HTTPPort = getPort()
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
