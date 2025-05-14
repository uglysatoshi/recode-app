package get

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.get"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		users, err := db.GetAllUsers()
		if err != nil {
			log.Error("failed to retrieve users", slog.Any("err", err))
			render.JSON(w, r, responce.Error("internal error"))
			return
		}

		render.JSON(w, r, map[string]any{
			"users": users,
		})
	}
}
