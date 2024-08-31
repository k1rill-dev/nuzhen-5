package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Lobby struct {
	gorm.Model
	id             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	name           string    `gorm:"type:varchar(255);unique;not null"`
	game           string    `gorm:"type:varchar(255);not null"`
	ruinerCount    int       `gorm:"default:0"`
	lobbyCount     int       `gorm:"default:0"`
	dateStart      time.Time `gorm:"default:null"`
	dateEnd        time.Time `gorm:"default:null"`
	additionalInfo string    `gorm:"type:text"`
	orgId          User      `gorm:"foreignKey:UserRefer"`
}

type LobbyStructure struct {
	gorm.Model
	id      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	lobbyId Lobby     `gorm:"foreignKey:LobbyRefer"`
	userId  *User     `gorm:"foreignKey:UserRefer"`
	chatId  Chat      `gorm:"foreignKey:ChatRefer"`
}
