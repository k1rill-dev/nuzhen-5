package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id             uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();" json:"id"`
	FirstName      string    `gorm:"type:varchar(255);" json:"first_name"`
	SeconderName   string    `gorm:"type:varchar(255);" json:"seconder_name"`
	ProfilePicture string    `gorm:"type:varchar(255);" json:"profile_picture"`
}
