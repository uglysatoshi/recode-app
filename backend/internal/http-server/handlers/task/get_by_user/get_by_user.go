package getbyuser

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
		const op = "handlers.task.get_by_user"

		// Получаем user_id из токена
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			log.Error("не удалось получить токен", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("unauthorized"))
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			log.Error("user_id отсутствует или неверного типа", slog.String("op", op))
			render.JSON(w, r, responce.Error("invalid token"))
			return
		}
		userID := uint(userIDFloat)

		// Получаем задачи по userID
		tasks, err := db.GetTasksByUserID(userID)
		if err != nil {
			log.Error("не удалось получить задачи", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to retrieve tasks"))
			return
		}

		render.JSON(w, r, responce.Data(tasks))
	}
}
