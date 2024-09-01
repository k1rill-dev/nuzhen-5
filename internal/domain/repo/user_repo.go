package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log/slog"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/models"
)

type UserRepo interface {
	SaveUser(user models.User) (models.User, error)
	GetUser(id uuid.UUID) (models.User, error)
}

type UserRepoImpl struct {
	DB  *gorm.DB
	log *slog.Logger
	cfg *config.Config
}

func NewUserRepoImpl(db *gorm.DB, log *slog.Logger, cfg *config.Config) *UserRepoImpl {
	return &UserRepoImpl{
		DB:  db,
		log: log,
		cfg: cfg,
	}
}

func (u *UserRepoImpl) SaveUser(user models.User) (models.User, error) {
	result := u.DB.Create(&user)
	return user, result.Error
}
func (u *UserRepoImpl) GetUser(id uuid.UUID) (models.User, error) {
	var user models.User
	result := u.DB.First(&user, id)
	return user, result.Error
}
