package main

import (
	"html/template"
	"log"
	"net/http"
)

type frontend struct{}

func (f *frontend) getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/front/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func (f *frontend) index(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/home.html")
	t.Execute(w, nil)
}

func (f *frontend) timeline(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/timeline.html")
	t.Execute(w, nil)
}

func (f *frontend) register(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/register.html")
	t.Execute(w, nil)
}

func (f *frontend) login(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/login.html")
	t.Execute(w, nil)
}
