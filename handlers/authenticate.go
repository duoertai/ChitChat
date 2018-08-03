package handlers

import (
	"net/http"
	"ChitChat/data"
	"fmt"
)

func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot%20parse%20form", 302)
		return
	}

	user, err := data.GetUserByEmail(request.PostFormValue("email"))
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot%20find%20user", 302)
		return
	}

	fmt.Println(user)
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			http.Redirect(writer, request, "/err?msg=Cannot%20create%20session", 302)
			return
		}

		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
