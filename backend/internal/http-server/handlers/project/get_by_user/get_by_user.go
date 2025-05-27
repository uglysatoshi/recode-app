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
		const op = "handlers.project.get_by_user"

		// Получение user_id из JWT
		_, claims, _ := jwtauth.FromContext(r.Context())
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			log.Error("user_id not found in token", slog.String("op", op))
			render.JSON(w, r, responce.Error("unauthorized"))
			return
		}
		userID := uint(userIDFloat)

		// Получение проектов
		projects, err := db.GetProjectsByUserID(userID)
		if err != nil {
			log.Error("failed to get projects", slog.String("op", op), slog.Any("err", err))
			render.JSON(w, r, responce.Error("failed to retrieve projects"))
			return
		}

		render.JSON(w, r, responce.Data(projects))
	}
}
