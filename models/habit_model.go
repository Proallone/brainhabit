package models

import (
	"brainhabit/models/common"

	"github.com/google/uuid"
)

type Habit struct {
	common.BaseModel
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"not null"`
	Description string
	Target      int `gorm:"not null"`
	Streak      int `gorm:"default:0"`
}
