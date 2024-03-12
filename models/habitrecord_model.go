package models

import (
	"brainhabit/models/common"

	"github.com/google/uuid"
)

type HabitRecord struct {
	common.BaseModel
	HabitID uuid.UUID `gorm:"type:uuid;not null"`
}
