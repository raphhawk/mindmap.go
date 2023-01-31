package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("../front-end/views/app.layout.gohtml", "../front-end/views/styles/main.css.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w)
	})
	fmt.Println("Starting front end server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
