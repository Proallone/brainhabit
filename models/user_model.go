package models

import "brainhabit/models/common"

type User struct {
	common.BaseModel
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
}
