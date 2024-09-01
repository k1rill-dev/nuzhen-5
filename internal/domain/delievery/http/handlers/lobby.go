package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
	"nuzhen-5-backend/internal/domain/repo"
)

type LobbyHandlers struct {
	cfg       *config.Config
	log       *slog.Logger
	lobbyRepo repo.LobbyRepo
}

func NewLobbyHandlers(cfg *config.Config, log *slog.Logger, lobbyRepo repo.LobbyRepo) *LobbyHandlers {
	return &LobbyHandlers{
		cfg:       cfg,
		log:       log,
		lobbyRepo: lobbyRepo,
	}
}

func (l *LobbyHandlers) CreateLobby(c *gin.Context) {
	var lobby models.Lobby
	if err := c.ShouldBindJSON(&lobby); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLobby, err := l.lobbyRepo.Save(lobby)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdLobby)
}

func (l *LobbyHandlers) GetLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	lobby, err := l.lobbyRepo.Get(lobbyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lobby)
}

func (l *LobbyHandlers) GetAllLobbiesFromUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	lobbies, err := l.lobbyRepo.GetAllLobbiesFromUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lobbies)
}

func (l *LobbyHandlers) AddUsersToLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("lobbyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var request struct {
		UserIDs []uuid.UUID `json:"user_ids"`
		ChatID  uuid.UUID   `json:"chat_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = l.lobbyRepo.AddUsersToLobby(lobbyID, request.UserIDs, request.ChatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "users added"})
}

func (l *LobbyHandlers) DeleteLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("lobbyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = l.lobbyRepo.DeleteLobby(lobbyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "lobby deleted"})
}

func (l *LobbyHandlers) RemoveUserFromLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("lobbyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	err = l.lobbyRepo.RemoveUserFromLobby(lobbyID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user removed"})
}

func (l *LobbyHandlers) UpdateLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("lobbyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var updates repo.LobbyUpdate
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLobby, err := l.lobbyRepo.UpdateLobby(lobbyID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLobby)
}

func (l *LobbyHandlers) GetAllUsersFromLobby(c *gin.Context) {
	lobbyID, err := uuid.Parse(c.Param("lobbyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	users, err := l.lobbyRepo.GetAllUsersFromLobby(lobbyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
