package save

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"backend/internal/models"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.save.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req models.User
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		// TODO: Добавить больше проверок
		if req.Username == "" || req.Email == "" || req.Password == "" {
			log.Warn("missing required fields")
			render.JSON(w, r, responce.Error("username, email and password are required"))
			return
		}

		id, err := db.SaveUser(req)
		if err != nil {
			log.Error("failed to save user", slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to save user"))
			return
		}

		log.Info("user saved", slog.Int64("id", int64(id)))

		render.JSON(w, r, map[string]any{
			"message": "user created",
			"id":      id,
		})
	}
}
