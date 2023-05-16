package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}
