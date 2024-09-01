package repo

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
	"time"
)

type LobbyRepo interface {
	Save(models.Lobby) (models.Lobby, error)
	Get(id uuid.UUID) (models.Lobby, error)
	GetAllLobbiesFromUserID(userID uuid.UUID) ([]models.Lobby, error)
	AddUsersToLobby(lobbyID uuid.UUID, userIDs []uuid.UUID, chatID uuid.UUID) error
	DeleteLobby(lobbyId uuid.UUID) error
	RemoveUserFromLobby(lobbyID uuid.UUID, userID uuid.UUID) error
	UpdateLobby(lobbyID uuid.UUID, updates LobbyUpdate) (models.Lobby, error)
	GetAllUsersFromLobby(lobbyID uuid.UUID) ([]models.User, error)
}

type LobbyUpdate struct {
	Name           *string    `json:"name,omitempty"`
	Game           *string    `json:"game,omitempty"`
	RuinerCount    *int       `json:"ruiner_count,omitempty"`
	LobbyCount     *int       `json:"lobby_count,omitempty"`
	DateStart      *time.Time `json:"date_start,omitempty"`
	DateEnd        *time.Time `json:"date_end,omitempty"`
	AdditionalInfo *string    `json:"additional_info,omitempty"`
	OrgID          *uuid.UUID `json:"org_id,omitempty"`
}

type LobbyRepoImpl struct {
	cfg *config.Config
	log *slog.Logger
	DB  *gorm.DB
}

func NewLobbyRepoImpl(cfg *config.Config, log *slog.Logger, db *gorm.DB) *LobbyRepoImpl {
	return &LobbyRepoImpl{
		cfg: cfg,
		log: log,
		DB:  db,
	}
}

func (l *LobbyRepoImpl) GetAllUsersFromLobby(lobbyID uuid.UUID) ([]models.User, error) {
	var users []models.User
	if err := l.DB.Where("lobby_id = ?", lobbyID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (l *LobbyRepoImpl) RemoveUserFromLobby(lobbyID uuid.UUID, userID uuid.UUID) error {
	var lobbyStructure models.LobbyStructure
	if err := l.DB.Where("lobby_id = ? AND user_id = ?", lobbyID, userID).First(&lobbyStructure).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("пользователь с ID %s не найден в лобби %s", userID, lobbyID)
		}
		return err
	}

	if err := l.DB.Delete(&lobbyStructure).Error; err != nil {
		return err
	}

	return nil

}

func (l *LobbyRepoImpl) Save(lobby models.Lobby) (models.Lobby, error) {
	result := l.DB.Create(&lobby)
	return lobby, result.Error
}
func (l *LobbyRepoImpl) Get(id uuid.UUID) (models.Lobby, error) {
	var lobby models.Lobby
	result := l.DB.First(&lobby, id)
	return lobby, result.Error
}

func (l *LobbyRepoImpl) GetAllLobbiesFromUserID(userID uuid.UUID) ([]models.Lobby, error) {
	var lobbies []models.Lobby
	result := l.DB.Where("user_id = ?", userID).Find(&lobbies)
	return lobbies, result.Error
}
func (l *LobbyRepoImpl) UpdateLobby(lobbyID uuid.UUID, updates LobbyUpdate) (models.Lobby, error) {
	var lobby models.Lobby
	if err := l.DB.First(&lobby, "id = ?", lobbyID).Error; err != nil {
		return models.Lobby{}, err
	}

	if err := l.DB.Model(&lobby).Updates(updates).Error; err != nil {
		return models.Lobby{}, err
	}
	return models.Lobby{}, nil

}

func (l *LobbyRepoImpl) AddUsersToLobby(lobbyID uuid.UUID, userIDs []uuid.UUID, chatID uuid.UUID) error {
	var lobby models.Lobby
	if err := l.DB.First(&lobby, "id = ?", lobbyID).Error; err != nil {
		return err
	}

	var existingUserCount int64
	if err := l.DB.Model(&models.LobbyStructure{}).Where("lobby_id = ?", lobbyID).Count(&existingUserCount).Error; err != nil {
		return err
	}

	if existingUserCount+int64(len(userIDs)) > int64(lobby.LobbyCount) {
		return fmt.Errorf("в лобби недостаточно места для добавления новых пользователей")
	}

	var newUsers []models.LobbyStructure

	for _, userID := range userIDs {
		found := false
		var existingUser models.LobbyStructure
		if err := l.DB.Where("lobby_id = ? AND user_id = ?", lobbyID, userID).First(&existingUser).Error; err == nil {
			found = true
		}

		if !found {
			newUsers = append(newUsers, models.LobbyStructure{
				LobbyId: lobbyID,
				UserId:  &userID,
				ChatId:  chatID,
			})
		}
	}

	if len(newUsers) > 0 {
		if err := l.DB.Create(&newUsers).Error; err != nil {
			return err
		}
	}

	return nil

}
func (l *LobbyRepoImpl) DeleteLobby(lobbyId uuid.UUID) error {
	result := l.DB.Delete(&models.Lobby{}, lobbyId)
	return result.Error
}
