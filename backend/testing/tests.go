package testing

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/http-server/handlers/user/login"
	"backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	cfg := &config.Config{
		DB: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
			SSLMode  string `yaml:"sslmode"`
		}(struct{ Host, Port, User, Password, Name, SSLMode string }{
			Host:     "localhost",
			Port:     "5432",
			User:     "test_user",
			Password: "test_pass",
			Name:     "test_db",
			SSLMode:  "disable",
		}),
	}
	store, err := database.New(cfg)
	if err != nil {
		t.Fatalf("failed to initialize storage: %v", err)
	}

	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	id, err := store.CreateUser(user)
	if err != nil {
		t.Fatalf("CreateUser returned error: %v", err)
	}
	if id == 0 {
		t.Fatalf("expected valid user ID, got 0")
	}
}

func TestLoginHandler(t *testing.T) {
	// Подготовка окружения
	cfg := &config.Config{
		DB: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Name     string `yaml:"name"`
			SSLMode  string `yaml:"sslmode"`
		}(struct{ Host, Port, User, Password, Name, SSLMode string }{
			Host:     "localhost",
			Port:     "5432",
			User:     "test_user",
			Password: "test_pass",
			Name:     "test_db",
			SSLMode:  "disable",
		}),
	}
	store, _ := database.New(cfg)

	// Создание тестового пользователя
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	store.CreateUser(models.User{
		Username: "user1",
		Email:    "user1@example.com",
		Password: string(hashedPassword),
	})

	// Тестируем login handler
	handler := login.New(nil, store)

	reqBody := strings.NewReader(`{"email":"user1@example.com", "password":"password123"}`)
	req := httptest.NewRequest(http.MethodPost, "/login", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", res.StatusCode)
	}
}
