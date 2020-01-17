package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles(templateFiles()...))

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
