package main

import (
	"encoding/json"
	"net/http"
)

func registerAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	var post *authPost
	resp := new(resp)
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		resp.errmsg = err.Error()
		encoder.Encode(resp)
		return
	}

	if post.username != "" && post.password != "" {
		user := new(user)
		user.username = post.username
		user.password = post.password
		users.Store(post.username, user)
	} else {
		resp.errmsg = "username or password is nil"
	}
	encoder.Encode(resp)
}
