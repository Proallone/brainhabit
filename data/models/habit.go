package models

import (
	"brainhabit/data/models/common"

	"github.com/google/uuid"
)

type Habit struct {
	common.BaseModel
	UserID      uuid.UUID `json:"userID" gorm:"type:uuid;not null"`
	Name        string    `json:"name" gorm:"not null;size:256"`
	Description string    `json:"description" gorm:"size:256"`
	Target      int       `json:"target" gorm:"not null"`
	Streak      int       `json:"streak" gorm:"default:0"`
	Records     []HabitRecord
}
