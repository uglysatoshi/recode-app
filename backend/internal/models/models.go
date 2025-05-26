package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model        // Включает стандартные поля ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique;not null"` // Уникальное имя пользователя (логин)
	Email      string `gorm:"unique;not null"` // Уникальный адрес электронной почты
	Password   string `gorm:"not null"`        // Пароль пользователя
	Role       string `gorm:"default:'user'"`  // Роль пользователя: user, admin и т.д.
}

type Task struct {
	gorm.Model            // Включает стандартные поля ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string    `gorm:"not null"` // Название задачи
	Description string    // Описание
	Status      string    `gorm:"default:'pending'"` // Статус задачи: pending, in-progress, done и т.д.
	Priority    string    `gorm:"default:'medium'"`  // Приоритет: low, medium, high
	Deadline    time.Time `gorm:"not null"`          // Дата и время окончания задачи
	UserID      uint      // Внешний ключ
	User        User      // Ассоциация с пользователем
	ProjectID   uint      // Внешний ключ
	Project     Project   // Ассоциация с проектом
}

type Project struct {
	gorm.Model         // Включает стандартные поля ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string `gorm:"not null"` // Название проекта
	Description string // Описание
	Tasks       []Task // Связь "один ко многим"
}
