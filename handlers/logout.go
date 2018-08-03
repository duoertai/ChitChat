package handlers

import (
	"net/http"
	"fmt"
	"ChitChat/data"
)

func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("failed to get cookie")
		session := data.Session{
			UUID: cookie.Value,
		}
		err = session.DeleteByUUID()
		if err != nil {
			fmt.Println("failed to delete session data")
		}
	}
	http.Redirect(writer, request, "/", 302)
}
