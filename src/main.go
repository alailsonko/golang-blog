package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "crud.com/src/docs"
)

type project struct {
	id    int
	techs []string
	url   string
}

var projects []project = []project{}

type profile struct {
	Name    string
	Hobbies []string
}

type obj struct {
	hello string
}

func postProject(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body := strings.NewReader(r.Body)
		var data obj
		decoder := json.NewDecoder(body).Decode(&data)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("this is a post", data)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	pro := profile{"alailson", []string{"snowboarding", "skateboarding"}}

	js, err := json.Marshal(pro)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.Unmarshal([]byte(js), &pro)
	fmt.Printf("%s\n", string(js))
	fmt.Printf("%s\n", pro.Hobbies)
	fmt.Printf("%s\n", pro.Name)
	w.Write(js)

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world")

}

// @title Swagger Example API
// @version 1.0
// @description golang crud.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host golang-crud.swagger.io
// @BasePath /v2
func main() {
	// r := chi.NewRouter()

	http.HandleFunc("/project/", postProject)
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/json/", foo)
	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL(getURL()),
	// ))

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "9090"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}

// func getURL() string {
// 	var urlSwagger = os.Getenv("URL_SWAGGER")
// 	// Set a default url if there is nothing in the environment
// 	if urlSwagger == "" {
// 		urlSwagger = "http://localhost:9090/swagger/doc.json"
// 		fmt.Println("INFO: No url environment variable detected, defaulting to " + urlSwagger)
// 	}
// 	return urlSwagger
// }
