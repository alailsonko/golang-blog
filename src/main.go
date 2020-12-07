package main

import (
	"fmt"
	_ "golang-CRUD/src/routers"
	"log"
	"os"
	"strconv"

	beego "github.com/astaxie/beego/server/web"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	portEnv := os.Getenv("PORT")
	fmt.Println("env", portEnv)
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatal(err)
	}
	beego.BConfig.Listen.HTTPPort = port
	beego.Run()
}
