package getbyuser

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Request struct {
	UserID uint `json:"user_id"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.project.get_by_user"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		projects, err := db.GetProjectsByUserID(req.UserID)
		if err != nil {
			log.Error("failed to get projects", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to retrieve projects"))
			return
		}

		render.JSON(w, r, responce.Data(projects))
	}
}
