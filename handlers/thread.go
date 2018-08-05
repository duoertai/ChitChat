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
