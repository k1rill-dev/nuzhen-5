package repo

import (
	"gorm.io/gorm"
	"log/slog"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
)

type ChatRepo interface {
	Save(chat models.Chat) (models.Chat, error)
}

type ChatRepoImpl struct {
	DB  *gorm.DB
	cfg *config.Config
	log *slog.Logger
}

func NewChatRepoImpl(db *gorm.DB, cfg *config.Config, log *slog.Logger) *ChatRepoImpl {
	return &ChatRepoImpl{
		DB:  db,
		cfg: cfg,
		log: log,
	}
}

func (r *ChatRepoImpl) Save(chat models.Chat) (models.Chat, error) {
	result := r.DB.Create(&chat)
	return chat, result.Error
}
