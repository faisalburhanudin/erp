package main

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

// User entity data
type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Role     string
}

// UserMgr for managing user data
type UserMgr struct {
	db *sqlx.DB
}

// ErrorWrongPassword indicate that password and hash is not match
var ErrorWrongPassword = errors.New("Password is not correct")

// GetByEmailPassword get user by email and password
func (mgr *UserMgr) GetByEmailPassword(email string, password string) (*User, error) {
	sql := "SELECT username, email, password, role FROM users WHERE email = ? LIMIT 1"

	user := User{}
	err := mgr.db.Get(&user, sql, email)
	if err != nil {
		log.Fatal(err)
	}

	isValid := CheckPasswordHash(password, user.Password)
	if isValid != true {
		return nil, ErrorWrongPassword
	}

	return &user, nil
}

// InsertUser insert user to databases
func (mgr *UserMgr) InsertUser(user *User) {
	user.Password = HashPassword(user.Password)
	sql := "INSERT INTO users (username, email, password, role) VALUES (:username, :email, :password, :role)"
	_, err := mgr.db.NamedExec(sql, user)
	if err != nil {
		log.Fatal(err)
	}
}

// GetUsers get list of user
func (mgr *UserMgr) GetUsers() []User {
	users := []User{}
	mgr.db.Select(&users, "SELECT username, email, password, role FROM users ORDER BY id desc")
	return users
}
