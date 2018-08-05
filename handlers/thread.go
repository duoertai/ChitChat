package handlers

import "net/http"

func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := checkSession(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}
