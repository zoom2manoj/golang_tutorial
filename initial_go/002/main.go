package main

import (
	"net/http"
	"html/template"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("initial_go/002/templates/*.gohtml"))
}

func main() {


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}


		tpl.ExecuteTemplate(writer, "layout.gohtml", data)
	})
	http.ListenAndServe(":8080", nil);
}
