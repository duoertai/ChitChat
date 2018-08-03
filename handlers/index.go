package handlers

import (
	"net/http"
	"ChitChat/data"
	"fmt"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.GetAllThreads()
	if err != nil {
		error_message(writer, request, err.Error())
	}

	_, err = checkSession(writer, request)
	if err != nil {
		fmt.Println(err)
		generateHTML(writer, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(writer, threads, "layout", "private.navbar", "index")
	}
}
