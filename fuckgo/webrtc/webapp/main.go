package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles(templateFiles()...))

var users sync.Map

func loginAPIHandler(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", registerAPIHandler).Methods("POST")
	router.HandleFunc("/api/login", loginAPIHandler).Methods("POST")

	router.HandleFunc("/index", baseAuth(indexHandler))
	router.HandleFunc("/login", loginHandler)
	panic(http.ListenAndServe(":8000", router))
}

func templateFiles() []string {
	templatePath := "./resource/"
	files, err := ioutil.ReadDir(templatePath)
	var paths []string
	if err == nil {
		for _, file := range files {
			fmt.Println(file.Name())
			paths = append(paths, templatePath+file.Name())
		}
	}
	return paths
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func baseAuth(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		authed := false
		if ok {
			value, ok := users.Load(username)
			if ok {
				user := value.(*user)
				authed = user.password == password
			}
		}

		if authed {
			fn(w, r)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
