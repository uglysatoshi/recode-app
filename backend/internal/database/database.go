package database

import (
	"backend/internal/config"
	"backend/internal/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.postgres.New"

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: failed to connect to DB: %w", op, err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("%s: migration failed: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(user models.User) (uint, error) {
	const op = "storage.postgres.SaveUser"

	result := s.db.Create(&user)
	if result.Error != nil {
		// Можно добавить проверку на дубликаты email, username, и вернуть кастомную ошибку
		return 0, fmt.Errorf("%s: failed to create user: %w", op, result.Error)
	}

	return user.ID, nil
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	const op = "storage.postgres.GetUserByEmail"

	var user models.User

	result := s.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%s: user not found", op)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%s: failed to get user: %w", op, result.Error)
	}

	return &user, nil
}

func (s *Storage) GetAllUsers() ([]models.User, error) {
	const op = "storage.postgres.GetAllUsers"

	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
