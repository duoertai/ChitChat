package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	// handle static resource
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	//mux.HandleFunc("/", )


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
