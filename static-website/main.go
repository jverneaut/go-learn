package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/contact.html"))
	t.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/contact", ContactHandler)

	http.ListenAndServe(":3000", r)
}
