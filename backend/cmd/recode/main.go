package main

import (
	"backend/internal/auth"
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/http-server/handlers/user/login"
	"backend/internal/http-server/handlers/user/register"
	"backend/internal/http-server/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jeffry-luqman/zlog"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	db, err := database.New(cfg)

	auth.InitJWT()

	if err != nil {
		log.Error("Failed to connect to database")
	}

	log.Info("Starting recode-app", slog.String("env", cfg.Env))

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // Добавляет уникальный ID для каждого запроса
	router.Use(middleware.RealIP)    // Определяет реальный IP-адрес клиента
	router.Use(middleware.Logger)    // Стандартное логгирование запросов
	router.Use(middleware.Recoverer) // Обработка паник и восстановление работы
	router.Use(middleware.URLFormat) // Парсинг URL-параметров
	router.Use(logger.New(log))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	log.Info("Starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	router.Route("/api/user", func(r chi.Router) {
		r.Post("/login", login.New(log, db))
		r.Post("/register", register.New(log, db))

		// Защищённые маршруты
		r.Group(func(protected chi.Router) {
			protected.Use(auth.Verifier())
			protected.Use(auth.Authenticator())
			// TODO: me, projects, tasks
		})
	})

	if err := srv.ListenAndServe(); err != nil {
		log.Error("Failed to start server")
	}

	log.Error("Server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		{
			zlog.HandlerOptions = &slog.HandlerOptions{Level: slog.LevelDebug}
			zlog.FmtDuration = []int{zlog.FgMagenta, zlog.FmtItalic}
			zlog.FmtPath = []int{zlog.FgHiCyan}
			log = zlog.New()
		}
	case envDev:
		{
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}

	case envProd:
		{
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

		}
	}
	return log
}
