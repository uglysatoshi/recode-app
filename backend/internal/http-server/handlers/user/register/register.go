package register

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"backend/internal/models"
	"encoding/json"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"strings"
)

type Request struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.register"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		req.Username = strings.TrimSpace(req.Username)

		if req.Email == "" || req.Username == "" || req.Password == "" {
			log.Warn("missing required fields", slog.String("op", op))
			render.JSON(w, r, responce.Error("username, email and password are required"))
			return
		}

		// Проверка: существует ли пользователь с таким email
		_, err := db.GetUserByEmail(req.Email)
		if err == nil {
			log.Warn("user already exists", slog.String("email", req.Email))
			render.JSON(w, r, responce.Error("user already exists"))
			return
		}

		// Хеширование пароля
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("failed to hash password", slog.Any("err", err))
			render.JSON(w, r, responce.Error("internal error"))
			return
		}

		user := models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: string(hashedPassword),
		}

		if _, err := db.CreateUser(user); err != nil {
			log.Error("failed to create user", slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to create user"))
			return
		}

		render.JSON(w, r, map[string]string{"message": "user registered successfully"})
	}
}
