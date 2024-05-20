package models

import (
	"brainhabit/data/models/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HabitRecord struct {
	common.BaseModel
	HabitID uuid.UUID `json:"habitID" gorm:"type:uuid;not null"`
}

func (hr *HabitRecord) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&Habit{}).Where("id = ?", hr.HabitID).Update("streak", gorm.Expr("streak + ?", 1))
	return
}

func (hr *HabitRecord) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&Habit{}).Where("id = ?", hr.HabitID).Update("streak", gorm.Expr("GREATEST(streak - ?, 0)", 1))
	return
}
