package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		w.Write([]byte("User ID: " + userID + "\n"))
	})

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		RequestAuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}
	log.Printf("Server has started %s\n", s.addr)
	return server.ListenAndServe()
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next.ServeHTTP
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequestAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
