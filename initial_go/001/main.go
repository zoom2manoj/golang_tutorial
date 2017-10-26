package main

import (
	"net/http"
	"fmt"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, Go Lang, how are you?", request.URL.Path)
	})
	http.ListenAndServe(":8080", nil);
}
