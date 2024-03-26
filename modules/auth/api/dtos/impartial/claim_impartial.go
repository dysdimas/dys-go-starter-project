package impartial

import "github.com/golang-jwt/jwt"

type ClaimImpartial struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}
