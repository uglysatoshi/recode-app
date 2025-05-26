package update

import (
	"backend/internal/database"
	"backend/internal/lib/api/responce"
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type Request struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	Deadline    time.Time `json:"deadline"`
	UserID      uint      `json:"user_id"`
}

func New(log *slog.Logger, db *database.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.task.update"

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("invalid request"))
			return
		}

		task, err := db.UpdateTask(req.ID, req.Title, req.Description, req.Status, req.Priority, req.Deadline, req.UserID)
		if err != nil {
			log.Error("failed to update task", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("could not update task"))
			return
		}

		render.JSON(w, r, task)
	}
}
