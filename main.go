package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo list",
		Todos: []Todo{
			{Item: "Install go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Master Go", Done: false},
		},
	}
	tmpl.Execute(w, data)

}

func main() {
	mux := http.NewServeMux()

	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	mux.HandleFunc("/todo", todo)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal("Error running server", err)
	}

}
