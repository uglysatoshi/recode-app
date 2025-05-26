package auth

import (
	"errors"
	"github.com/go-chi/jwtauth"
	"net/http"
)

type contextKey string

const userCtxKey = contextKey("user_id")

func Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(TokenAuth)
}

func Authenticator() func(http.Handler) http.Handler {
	return jwtauth.Authenticator
}

func UserIDFromContext(r *http.Request) (uint, error) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	idRaw, ok := claims["user_id"]
	if !ok {
		return 0, errors.New("user_id not found in token")
	}

	idFloat, ok := idRaw.(float64)
	if !ok {
		return 0, errors.New("invalid user_id type")
	}

	return uint(idFloat), nil
}
