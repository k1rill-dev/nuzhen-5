package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ChatId string    `gorm:"varchar(255);" json:"chat_id"`
	UserId uuid.UUID `json:"user"`
	User   User
}
