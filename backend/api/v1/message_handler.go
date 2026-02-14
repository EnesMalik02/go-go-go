package v1

import (
	"gin-tutorial/internal/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service *message.Service
}

func NewMessageHandler(service *message.Service) *MessageHandler {
	return &MessageHandler{service: service}
}

func (h *MessageHandler) RegisterRoutes(r *gin.RouterGroup) {
	messages := r.Group("/messages")
	{
		messages.GET("", h.GetMessages)
		messages.POST("", h.CreateMessage)
		messages.PUT("/:id", h.UpdateMessage)
		messages.DELETE("/:id", h.DeleteMessage)
	}
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	messages := h.service.GetMessages()
	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var newMsg message.Message
	if err := c.ShouldBindJSON(&newMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMsg := h.service.CreateMessage(newMsg)
	c.JSON(http.StatusCreated, createdMsg)
}

func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := h.service.UpdateMessage(id, updateData.Content)
	if msg == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	success := h.service.DeleteMessage(id)
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
