package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// User entity data
type User struct {
	username string
	email    string
	password string
	role     string
}

// Model is Database instance
type Model struct {
	db *sqlx.DB
}

// NewModel create new instance Model
func NewModel(dsn string) Model {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return Model{db: db}
}

// InsertUser insert user to databases
func (m *Model) InsertUser(user *User) {
	sql := "INSERT INTO user (username, email, password, role) VALUES (:username, :email, :password, :role)"
	_, err := m.db.NamedExec(sql, user)
	if err != nil {
		log.Fatal(err)
	}
}
