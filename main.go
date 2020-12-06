package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

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

func locateFiles(path string) string {
	cwd, _ := os.Getwd()
	location := filepath.Join(cwd, path)
	return location

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
		filePath := "./src/login.gtpl"
		t, _ := template.ParseFiles(locateFiles(filePath))
		err := t.Execute(w, token)
		if err != nil {
			fmt.Println("err:", err)
		}
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		fmt.Println("my token: ", token)
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles(locateFiles("./src/upload.gtpl"))
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(locateFiles("./src/files/")+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

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
