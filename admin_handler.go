package main

import (
	"html/template"
	"log"
	"net/http"
)

type admin struct{}

func (a *admin) getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/admin/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func (a *admin) index(w http.ResponseWriter, r *http.Request) {
	t := a.getTemplate("templates/admin/home.html")
	t.Execute(w, nil)
}

func (a *admin) employeeNew(w http.ResponseWriter, r *http.Request) {
	t := a.getTemplate("templates/admin/employee_new.html")
	t.Execute(w, nil)
}
