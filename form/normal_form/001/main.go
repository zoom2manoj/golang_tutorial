package main

import (
	"net/http"
	"html/template"
)

type ContactDetails struct {
	Email string
	Subject string
	Message string
}



func main() {
	tmp := template.Must(template.ParseFiles("form/normal_form/001/web/index.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		if request.Method != http.MethodPost {
			tmp.Execute(writer, nil)
			return
		}

		details := ContactDetails{
			Email:"manoj.kb2011@gmail.com",
			Subject:"Form submmition",
			Message:"Message body here",
		}

		_ = details

		tmp.Execute(writer, struct {
			Success bool
		}{true})

	})
	http.ListenAndServe(":8080", nil);
}




