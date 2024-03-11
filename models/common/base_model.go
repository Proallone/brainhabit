package common

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` //* override default ID with UUID
}
