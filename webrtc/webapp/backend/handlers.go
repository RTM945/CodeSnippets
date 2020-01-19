package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	var post *authPost
	var err error
	resp := new(resp)
	err = json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		resp.errmsg = err.Error()
		encoder.Encode(resp)
		return
	}

	if post.username != "" && post.password != "" {
		err = register(post.username, post.password)
		if err != nil {
			resp.errmsg = err.Error()
		}
	} else {
		resp.errmsg = "username or password is nil"
	}
	encoder.Encode(resp)
}

func pairHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderTemplate(w, "login")
		return
	}
	redirect := "/login"
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username != "" && password != "" {
		if auth(username, password) {
			setSession(username, w)
			redirect = "/"
		}
	}
	http.Redirect(w, r, redirect, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}
