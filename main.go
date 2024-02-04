package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const localhostPort = ":8080"

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("todo.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Clean", Done: false},
				{Title: "Work", Done: true},
				{Title: "Work", Done: true},
			},
		}
		err := tmpl.Execute(w, data)

		if err != nil {
			fmt.Println("Error parsing template page")
		}
	})

	// handles statics filesimages, css, js
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(localhostPort, nil)

	if err != nil {
		fmt.Println("Error connecting to localhost", localhostPort)
	}

}
