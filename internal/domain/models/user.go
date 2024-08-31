package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	firstName      string    `gorm:"type:varchar(255);"`
	seconderName   string    `gorm:"type:varchar(255);"`
	profilePicture string    `gorm:"type:varchar(255);"`
}
