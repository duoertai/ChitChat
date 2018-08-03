package main

import (
	"net/http"
	"time"
	"ChitChat/handlers"
)

func main() {
	mux := http.NewServeMux()

	// handle static resource
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", handlers.Index)
	// error
	mux.HandleFunc("/err", handlers.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/signup_account", handlers.SignupAccount)
	mux.HandleFunc("/authenticate", handlers.Authenticate)

	// create a server
	server := &http.Server{
		Addr: config.Address,
		Handler: mux,
		ReadTimeout: time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
