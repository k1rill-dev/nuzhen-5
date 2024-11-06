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
	lobbyHandlers *handlers.LobbyHandlers
}

func NewHTTPServer(cfg *config.Config, log *slog.Logger, userHandlers *handlers.UserHandlers,
	chatHandlers *handlers.ChatHandlers, lobbyHandlers *handlers.LobbyHandlers) *HTTPServer {
	return &HTTPServer{
		cfg:           cfg,
		log:           log,
		userHandlers:  userHandlers,
		chatHandlers:  chatHandlers,
		lobbyHandlers: lobbyHandlers,
	}
}

func (h *HTTPServer) Run() {
	router := gin.Default()

	router.GET("/ping", h.userHandlers.Ping)
	router.POST("/users", h.userHandlers.CreateUser)
	router.GET("/users/:userId/details", h.userHandlers.GetUser)

	router.POST("/chats", h.chatHandlers.CreateChat)

	router.POST("/lobbies", h.lobbyHandlers.CreateLobby)
	router.GET("/lobbies/getInfo/:lobbyId", h.lobbyHandlers.GetLobbyInfo)
	router.GET("/lobbies/:lobbyId", h.lobbyHandlers.GetLobby)
	router.GET("/users/:userId/lobbies", h.lobbyHandlers.GetAllLobbiesFromUser)
	router.POST("/lobbies/:lobbyId/users", h.lobbyHandlers.AddUsersToLobby)
	router.DELETE("/lobbies/:lobbyId", h.lobbyHandlers.DeleteLobby)
	router.DELETE("/lobbies/:lobbyId/users/:userId", h.lobbyHandlers.RemoveUserFromLobby)
	router.PUT("/lobbies/:lobbyId", h.lobbyHandlers.UpdateLobby)
	router.GET("/lobbies/:lobbyId/users", h.lobbyHandlers.GetAllUsersFromLobby)

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
