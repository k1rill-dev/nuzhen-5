package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
	"nuzhen-5-backend/internal/domain/repo"
	"nuzhen-5-backend/pkg/lib/log"
)

type UserHandlers struct {
	cfg      *config.Config
	log      *slog.Logger
	userRepo repo.UserRepo
}

func NewUserHandlers(cfg *config.Config, log *slog.Logger, userRepo repo.UserRepo) *UserHandlers {
	return &UserHandlers{
		cfg:      cfg,
		log:      log,
		userRepo: userRepo,
	}
}

func (h *UserHandlers) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
	var user models.User
	var err error
	if err := c.BindJSON(&user); err != nil {
		log.Err(err)
	}
	user, err = h.userRepo.SaveUser(user)
	if err != nil {
		log.Err(err)
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func (h *UserHandlers) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userRepo.GetUser(uuid.MustParse(id))
	if err != nil {
		log.Err(err)
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}
