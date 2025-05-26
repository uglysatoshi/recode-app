package me

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.me"

		_, claims, _ := jwtauth.FromContext(r.Context())

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			log.Error("invalid user_id claim", slog.String("op", op))
			render.JSON(w, r, responce.Error("unauthorized"))
			return
		}
		userID := uint(userIDFloat)

		user, err := db.GetUserByID(userID)
		if err != nil {
			log.Error("failed to get user", slog.Any("err", err))
			render.JSON(w, r, responce.Error("user not found"))
			return
		}

		render.JSON(w, r, responce.Data(map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		}))
	}
}
