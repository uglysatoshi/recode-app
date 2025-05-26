package update

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Request struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.project.update"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		project, err := db.UpdateProject(req.ID, req.Title, req.Description)
		if err != nil {
			log.Error("failed to update project", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("could not update project"))
			return
		}

		render.JSON(w, r, project)
	}
}
