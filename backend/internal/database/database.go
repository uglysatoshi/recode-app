package database

import (
	"backend/internal/config"
	"backend/internal/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
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

	if err := db.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.Project{},
	); err != nil {
		return nil, fmt.Errorf("%s: migration failed: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(user models.User) (uint, error) {
	const op = "storage.postgres.SaveUser"

	result := s.db.Create(&user)
	if result.Error != nil {
		// TODO: Можно добавить проверку на дубликаты email, username
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

func (s *Storage) CreateProject(title, description string) (*models.Project, error) {
	project := &models.Project{
		Title:       title,
		Description: description,
	}
	if err := s.db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Storage) CreateTask(title, desc, status, priority string, deadline time.Time, projectID, userID uint) (*models.Task, error) {
	task := &models.Task{
		Title:       title,
		Description: desc,
		Status:      status,
		Priority:    priority,
		Deadline:    deadline,
		ProjectID:   projectID,
		UserID:      userID,
	}
	if err := s.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Storage) UpdateProject(id uint, title, description string) (*models.Project, error) {
	var project models.Project
	if err := s.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	project.Title = title
	project.Description = description
	if err := s.db.Save(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *Storage) UpdateTask(id uint, title, desc, status, priority string, deadline time.Time, userID uint) (*models.Task, error) {
	var task models.Task
	if err := s.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	task.Title = title
	task.Description = desc
	task.Status = status
	task.Priority = priority
	task.Deadline = deadline
	task.UserID = userID
	if err := s.db.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *Storage) DeleteProject(id uint) error {
	return s.db.Delete(&models.Project{}, id).Error
}

func (s *Storage) DeleteTask(id uint) error {
	return s.db.Delete(&models.Task{}, id).Error
}

func (s *Storage) GetTasksByProjectID(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Where("project_id = ?", projectID).Find(&tasks).Error
	return tasks, err
}

func (s *Storage) GetProjectsByUserID(userID uint) ([]models.Project, error) {
	var projects []models.Project
	err := s.db.Joins("JOIN tasks ON tasks.project_id = projects.id").
		Where("tasks.user_id = ?", userID).
		Group("projects.id").
		Find(&projects).Error
	return projects, err
}

func (s *Storage) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (s *Storage) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
