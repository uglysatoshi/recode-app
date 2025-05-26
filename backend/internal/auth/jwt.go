package auth

import (
	"github.com/go-chi/jwtauth"
	"os"
)

var TokenAuth *jwtauth.JWTAuth

func InitJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set in environment")
	}
	TokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}
