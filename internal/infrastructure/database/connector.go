package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nuzhen-5-backend/config"
)

type IDataBase interface {
	Connect() (*gorm.DB, error)
	CreateTables(db *gorm.DB, models ...interface{}) error
}

type Postgres struct {
	IDataBase
	cfg *config.Config
}

func NewPostgresConnection(config *config.Config) *Postgres {
	return &Postgres{
		cfg: config,
	}
}

func (p *Postgres) Connect() (*gorm.DB, error) {
	// postgres://pg:pass@localhost:5432/crud
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", p.cfg.DatabaseUser,
		p.cfg.DatabasePassword, p.cfg.DatabaseHost, p.cfg.DatabasePort, p.cfg.DatabaseName)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (p *Postgres) CreateTables(db *gorm.DB, models ...interface{}) error {
	for _, model := range models {
		err := db.Migrator().CreateTable(&model)
		if err != nil {
			return err
		}
	}
	return nil
}
