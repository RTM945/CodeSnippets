package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/", authWapper(indexHandler))
	router.HandleFunc("/pair", pairHandler)
	panic(http.ListenAndServe(":8000", router))
}

func authWapper(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
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
			http.Redirect(w, r, "/pair", http.StatusFound)
		}
	}
}
