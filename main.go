package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		title := "Азбука"
		author := "Иван Федоров"
		fmt.Fprintf(w, `{"author":%s , "title": %s}`, author, title)
	}).Methods("POST")

	http.ListenAndServe(":8082", r)
}
