package v1

import (
	"gin-tutorial/internal/item"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ItemHandler handles HTTP requests for items
type ItemHandler struct {
	service *item.Service
}

// NewItemHandler creates a new ItemHandler
func NewItemHandler(service *item.Service) *ItemHandler {
	return &ItemHandler{service: service}
}

// RegisterRoutes registers the item routes
func (h *ItemHandler) RegisterRoutes(r *gin.RouterGroup) {
	items := r.Group("/items")
	{
		items.GET("", h.GetItems)
		items.GET("/:id", h.GetItem)
		items.POST("", h.CreateItem)
		items.PUT("/:id", h.UpdateItem)
		items.DELETE("/:id", h.DeleteItem)
	}
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	skipStr := c.DefaultQuery("skip", "0")
	limitStr := c.DefaultQuery("limit", "10")

	skip, _ := strconv.Atoi(skipStr)
	limit, _ := strconv.Atoi(limitStr)

	items := h.service.GetItems(skip, limit)
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	item, err := h.service.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var newItem item.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdItem := h.service.CreateItem(newItem)
	c.JSON(http.StatusCreated, createdItem)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var updatedItem item.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.UpdateItem(id, updatedItem)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
