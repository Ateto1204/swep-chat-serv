package delivery

import (
	"net/http"

	"github.com/Ateto1204/swep-chat-serv/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatUseCase usecase.ChatUseCase
}

func NewChatHandler(chatUseCase usecase.ChatUseCase) *ChatHandler {
	return &ChatHandler{chatUseCase}
}

func (h *ChatHandler) SaveChat(c *gin.Context) {
	type Input struct {
		Name    string   `json:"name"`
		Members []string `json:"members"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chat, err := h.chatUseCase.SaveChat(input.Name, input.Members)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chat)
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	type Input struct {
		ID string `json:"id"`
	}
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chat, err := h.chatUseCase.GetChat(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chat)
}
