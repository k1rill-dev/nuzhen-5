package main

import (
	"log/slog"
	cfg "nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/delievery/http"
	"nuzhen-5-backend/internal/domain/delievery/http/handlers"
	"nuzhen-5-backend/internal/domain/models"
	"nuzhen-5-backend/internal/domain/repo"
	"nuzhen-5-backend/internal/infrastructure/database"
	"nuzhen-5-backend/internal/infrastructure/di"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := cfg.MustLoadConfig("./config/config.yaml")
	log := setupLogger(config.Env)
	log.Info("config loaded successfully!")
	container := di.NewContainer()
	connector := database.NewPostgresConnection(config)
	postgresDb, err := connector.Connect()
	userRepo := repo.NewUserRepoImpl(postgresDb, log, config)
	chatRepo := repo.NewChatRepoImpl(postgresDb, config, log)
	lobbyRepo := repo.NewLobbyRepoImpl(config, log, postgresDb)
	userHandler := handlers.NewUserHandlers(config, log, userRepo)
	chatHandler := handlers.NewChatHandlers(config, log, chatRepo)
	lobbyHandler := handlers.NewLobbyHandlers(config, log, lobbyRepo)
	httpServer := http.NewHTTPServer(config, log, userHandler, chatHandler, lobbyHandler)
	if err != nil {
		panic(err)
	}

	err = connector.CreateTables(postgresDb, &models.User{}, &models.Chat{}, &models.Lobby{}, &models.LobbyStructure{})
	if err != nil {
		log.Error("Failed to create tables, because they exists")
	}

	log.Info("postgres connected successfully!")

	// register dependencies
	container.Register(log)
	container.Register(config)
	container.Register(container)
	container.Register(postgresDb)
	log.Info("dependencies registry successfully!")

	httpServer.Run()

	log.Info("shutdown")

}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
