package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Lobby struct {
	gorm.Model
	Id             uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name           string    `gorm:"type:varchar(255);unique;not null" json:"name"`
	Game           string    `gorm:"type:varchar(255);not null" json:"game"`
	RuinerCount    int       `gorm:"default:0" json:"ruiner_count"`
	LobbyCount     int       `gorm:"default:0" json:"lobby_count"`
	DateStart      time.Time `gorm:"default:null" json:"date_start"`
	DateEnd        time.Time `gorm:"default:null" json:"date_end"`
	AdditionalInfo string    `gorm:"type:text" json:"additional_info"`
	OrgID          uuid.UUID `json:"org_id"`
	Org            User
}

type LobbyStructure struct {
	gorm.Model
	Id      uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	LobbyId uuid.UUID  `json:"lobby_id"`
	UserId  *uuid.UUID `json:"user_id"`
	ChatId  uuid.UUID  `json:"chat_id"`
	Lobby   Lobby
	User    User
	Chat    Chat
}
