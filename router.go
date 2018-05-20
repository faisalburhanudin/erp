package main

import (
	"html/template"
	"log"
	"net/http"
)

func getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/front/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func index(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("templates/front/home.html")
	t.Execute(w, nil)
}

func timeline(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("templates/front/timeline.html")
	t.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("templates/front/register.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t := getTemplate("templates/front/login.html")
	t.Execute(w, nil)
}
