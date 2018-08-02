package handlers

import "net/http"

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := checkSession(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
