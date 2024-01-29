package main

import (
	"fmt"
	"net/http"
)

type task struct {
	todo   string
	isDone bool
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Bar")
}

func todo(w http.ResponseWriter, req *http.Request) {
	t := task{todo: "clean", isDone: false}
	fmt.Fprint(w, t)
}

func main() {

	http.HandleFunc("/foo", foo)
	http.HandleFunc("/todo", todo)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("error connecting to localhost")
	}
}
