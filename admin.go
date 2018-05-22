package main

import (
	"html/template"
	"log"
	"net/http"
)

// Admin handler
type Admin struct {
	userMgr UserMgr
}

func (a *Admin) getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(
		"templates/admin/base.html",
		templateName,
	)
	if err != nil {
		log.Fatalf("Error load template %v", err)
	}

	return t
}

func (a *Admin) index(w http.ResponseWriter, r *http.Request) {
	t := a.getTemplate("templates/admin/home.html")
	t.Execute(w, nil)
}

func (a *Admin) employeeNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		role := r.FormValue("role")

		user := User{
			Username: username,
			Email:    email,
			Password: password,
			Role:     role,
		}
		a.userMgr.InsertUser(&user)

		http.Redirect(w, r, "/admin/employee/list", 301)
	} else {

	}
	t := a.getTemplate("templates/admin/employee/new.html")
	t.Execute(w, nil)
}

// EmployeeListData for template
type EmployeeListData struct {
	Users []User
}

func (a *Admin) employeeList(w http.ResponseWriter, r *http.Request) {
	users := a.userMgr.GetUsers()

	data := EmployeeListData{
		Users: users,
	}

	t := a.getTemplate("templates/admin/employee/list.html")
	t.Execute(w, data)
}
