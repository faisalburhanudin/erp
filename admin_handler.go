package main

import (
	"html/template"
	"log"
	"net/http"
)

type admin struct{}

func (b *admin) getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/admin/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func (b *admin) index(w http.ResponseWriter, r *http.Request) {
	t := b.getTemplate("templates/admin/home.html")
	t.Execute(w, nil)
}
