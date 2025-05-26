package create

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.project.create"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		project, err := db.CreateProject(req.Title, req.Description)
		if err != nil {
			log.Error("failed to create project", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("could not create project"))
			return
		}

		render.JSON(w, r, project)
	}
}
