package handlers

import (
	"net/http"
	"ChitChat/data"
)

func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot%20parse%20form", 302)
		return
	}

	user := data.User{
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	if err = user.CreateUser(); err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot%20create%20user", 302)
		return
	}

	http.Redirect(writer, request, "/login", 302)
}
