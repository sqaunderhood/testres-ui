package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseGlob(filepath.Join(*staticDir, "templates/*.html")))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
