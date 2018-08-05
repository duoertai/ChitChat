package handlers

import (
	"net/http"
	"ChitChat/data"
)

func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := checkSession(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

func CreateThread(writer http.ResponseWriter, request *http.Request) {
	session, err := checkSession(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20parse%20form", 302)
			return
		}

		user, err := session.GetUserFromSession()
		if err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20find%20user", 302)
			return
		}

		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20create%20thread", 302)
			return
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func ReadThread(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	uuid := values.Get("id")
	thread, err := data.GetThreadByUUID(uuid)
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot%20create%20thread", 302)
		return
	} else {
		_, err := checkSession(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}
