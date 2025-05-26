package login

import (
	"backend/internal/auth"
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"time"
)

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.login"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		user, err := db.GetUserByEmail(req.Email)
		if err != nil {
			render.JSON(w, r, responce.Error("invalid email or password"))
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			render.JSON(w, r, responce.Error("invalid email or password"))
			return
		}

		// Выдача JWT токена
		_, tokenString, _ := auth.TokenAuth.Encode(map[string]interface{}{
			"user_id": user.ID,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})

		render.JSON(w, r, map[string]string{"token": tokenString})
	}
}
