package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("web/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "message", nil) // Render only the "message" fragment
	})

	http.ListenAndServe(":8080", nil)
}
