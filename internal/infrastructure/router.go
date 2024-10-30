package infrastructure

import (
	"github.com/Ateto1204/swep-chat-serv/internal/delivery"
	"github.com/Ateto1204/swep-chat-serv/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(chatUseCase usecase.ChatUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	handler := delivery.NewChatHandler(chatUseCase)

	router.POST("/api/chat", handler.SaveChat)
	router.POST("/api/chat/id", handler.GetChat)

	return router
}
