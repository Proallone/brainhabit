package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	DisplayName  string    `json:"display_name"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
}
