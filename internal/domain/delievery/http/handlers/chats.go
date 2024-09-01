package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
	"nuzhen-5-backend/internal/domain/repo"
	"nuzhen-5-backend/pkg/lib/log"
)

type ChatHandlers struct {
	cfg      *config.Config
	log      *slog.Logger
	chatRepo repo.ChatRepo
}

func NewChatHandlers(cfg *config.Config, log *slog.Logger, chatRepo repo.ChatRepo) *ChatHandlers {
	return &ChatHandlers{
		cfg:      cfg,
		log:      log,
		chatRepo: chatRepo,
	}
}

func (h *ChatHandlers) CreateChat(c *gin.Context) {
	var chat models.Chat
	var err error
	if err := c.ShouldBindJSON(&chat); err != nil {
		log.Err(err)
	}
	chat, err = h.chatRepo.Save(chat)
	if err != nil {
		log.Err(err)
	}
	c.JSON(http.StatusCreated, chat)
}
