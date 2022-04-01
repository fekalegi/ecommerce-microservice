package models

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	AuthID   string `json:"auth_id"`
	jwt.StandardClaims
}
