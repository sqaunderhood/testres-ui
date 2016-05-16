package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("static/templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
