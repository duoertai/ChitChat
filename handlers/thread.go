package handlers

import (
	"net/http"
	"ChitChat/data"
	"fmt"
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

func PostThread(writer http.ResponseWriter, request *http.Request) {
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

		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.GetThreadByUUID(uuid)
		if err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20find%20thread", 302)
			return
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20create%20post", 302)
			return
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}

