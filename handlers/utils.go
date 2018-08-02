package handlers

import (
	"net/http"
	"strings"
	"ChitChat/data"
	"errors"
	"fmt"
	"html/template"
)

func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func checkSession(writer http.ResponseWriter, request *http.Request) (session data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		return session, err
	}

	session = data.Session{
		UUID: cookie.Value,
	}

	valid, err := session.CheckSession()
	if err != nil {
		return session, err
	}

	if valid {
		return session, err
	}

	return session, errors.New("invalid session")
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("../templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
