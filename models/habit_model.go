package models

import (
	"brainhabit/models/common"

	"github.com/google/uuid"
)

type Habit struct {
	common.BaseModel
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"not null;size:256"`
	Description string    `gorm:"size:256"`
	Target      int       `gorm:"not null"`
	Streak      int       `gorm:"default:0"`
}
