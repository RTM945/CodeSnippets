package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/", authMiddleware(indexHandler))
	router.HandleFunc("/pair", pairHandler)
	panic(http.ListenAndServe(":8000", router))
}

func authMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := getUserName(r)
		if username != "" {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/pair", http.StatusFound)
		}
	})
}
