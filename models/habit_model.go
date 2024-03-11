package models

import "brainhabit/models/common"

type Habit struct {
	common.BaseModel
	UserID      uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Target      int `gorm:"not null"`
	Streak      int `gorm:"default:0"`
}
