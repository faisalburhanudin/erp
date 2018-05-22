package main

import (
	"html/template"
	"log"
	"net/http"
)

// Frontend handler
type Frontend struct {
	userMgr UserMgr
}

func (f *Frontend) getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/front/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func (f *Frontend) index(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/home.html")
	t.Execute(w, nil)
}

func (f *Frontend) timeline(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/timeline.html")
	t.Execute(w, nil)
}

func (f *Frontend) register(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/register.html")
	t.Execute(w, nil)
}

func (f *Frontend) login(w http.ResponseWriter, r *http.Request) {
	t := f.getTemplate("templates/front/login.html")
	t.Execute(w, nil)
}
