package models

import "github.com/leandropaes/goapi/lib"

// Seller table users
type Seller struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required"`
}

// Sellers array from Seller
type Sellers []Seller

// SellerModel load table users from db
var SellerModel = lib.Sess.Collection("sellers")
