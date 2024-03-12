package models

import "brainhabit/models/common"

type User struct {
	common.BaseModel
	Name         string `json:"name" gorm:"size:256"`
	Surname      string `json:"surname" gorm:"size:256"`
	DisplayName  string `json:"display_name" gorm:"size:256"`
	Email        string `json:"email" gorm:"uniqueIndex;not null;size:256"`
	PasswordHash string `json:"-" gorm:"not null;size:256"`
	Habits       []Habit
}
