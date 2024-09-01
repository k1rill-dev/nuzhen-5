package http

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"nuzhen-5-backend/config"
	"nuzhen-5-backend/internal/domain/delievery/http/handlers"
)

type HTTPServer struct {
	cfg           *config.Config
	log           *slog.Logger
	userHandlers  *handlers.UserHandlers
	chatHandlers  *handlers.ChatHandlers
	lobbyHandlers *handlers.LobbyHandlers // Добавляем поле для LobbyHandlers
}

func NewHTTPServer(cfg *config.Config, log *slog.Logger, userHandlers *handlers.UserHandlers,
	chatHandlers *handlers.ChatHandlers, lobbyHandlers *handlers.LobbyHandlers) *HTTPServer { // Добавляем lobbyHandlers в параметры
	return &HTTPServer{
		cfg:           cfg,
		log:           log,
		userHandlers:  userHandlers,
		chatHandlers:  chatHandlers,
		lobbyHandlers: lobbyHandlers, // Инициализируем lobbyHandlers
	}
}

func (h *HTTPServer) Run() {
	router := gin.Default()

	// Маршруты для UserHandlers
	router.GET("/ping", h.userHandlers.Ping)
	router.POST("/users", h.userHandlers.CreateUser)
	router.GET("/users/:id", h.userHandlers.GetUser)

	// Маршруты для ChatHandlers
	router.GET("/chats", h.chatHandlers.CreateChat)

	// Маршруты для LobbyHandlers
	router.POST("/lobbies", h.lobbyHandlers.CreateLobby)
	router.GET("/lobbies/:id", h.lobbyHandlers.GetLobby)
	router.GET("/users/:userID/lobbies", h.lobbyHandlers.GetAllLobbiesFromUser)
	router.POST("/lobbies/:lobbyID/users", h.lobbyHandlers.AddUsersToLobby)
	router.DELETE("/lobbies/:lobbyID", h.lobbyHandlers.DeleteLobby)
	router.DELETE("/lobbies/:lobbyID/users/:userID", h.lobbyHandlers.RemoveUserFromLobby)
	router.PUT("/lobbies/:lobbyID", h.lobbyHandlers.UpdateLobby)
	router.GET("/lobbies/:lobbyID/users", h.lobbyHandlers.GetAllUsersFromLobby)

	// Запуск сервера
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
