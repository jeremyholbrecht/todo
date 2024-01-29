package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
				{Title: "Meditate", Done: true},
			},
		}
		err := tmpl.Execute(w, data)

		if err != nil {
			fmt.Println("Error parsing template page")
		}
	})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error connecting to localhost:8080")
	}

}
