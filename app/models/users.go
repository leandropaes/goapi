package models

import "github.com/leandropaes/goapi/lib"

// User table users
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name" validate:"required"`
	Email string `db:"email" json:"email" validate:"email"`
	Password string `db:"password" json:"-" validate:"required"`
}

// Users array from User
type Users []User

// UserModel load table users from db
var UserModel = lib.Sess.Collection("users")
