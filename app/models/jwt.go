package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
