package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	id     uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	chatId string    `gorm:"varchar(255);"`
	userId User      `gorm:"foreignKey:UserRefer"`
}
