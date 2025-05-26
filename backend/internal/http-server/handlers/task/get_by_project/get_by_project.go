package getbyproject

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Request struct {
	ProjectID uint `json:"project_id"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.task.get_by_project"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		tasks, err := db.GetTasksByProjectID(req.ProjectID)
		if err != nil {
			log.Error("failed to get tasks", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to retrieve tasks"))
			return
		}

		render.JSON(w, r, responce.Data(tasks))
	}
}
