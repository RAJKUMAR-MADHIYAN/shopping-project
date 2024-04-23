package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"shops/models"
	"shops/services"

	"github.com/gin-gonic/gin"
)

type ItemHandler interface {
	ListItems(c *gin.Context)
	CreateItem(c *gin.Context)
	FindItem(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}
type itemHandler struct {
	s *services.ItemService
}

func NewItemHandler(s services.ItemService) ItemHandler {
	if s == nil {
		log.Fatal("Failed to initialize item handler, service is nil")
		return nil
	}
	var p = itemHandler{s: &s}
	return &p
}

func (h *itemHandler) GetItemService() services.ItemService {
	if h.s == nil {
		log.Fatal("Failed to get item service, it is nil")
		return nil
	}

	return *h.s
}
func (h *itemHandler) ListItems(c *gin.Context) {
	s := h.GetItemService()
	items, err := s.ListItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *itemHandler) FindItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bad id value: %s", err)})
		return
	}
	s := h.GetItemService()
	item, found, err := s.FindItem(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}
	if found {
		c.JSON(http.StatusOK, gin.H{"data": item})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"data": nil})
}

func (h *itemHandler) CreateItem(c *gin.Context) {
	var input models.CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s := h.GetItemService()
	item, err := s.CreateItem(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}
func (h *itemHandler) UpdateItem(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bad id value: %s", e)})
		return
	}
	// Validate input
	var input models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := h.GetItemService()
	item, err := s.UpdateItem(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}
func (h *itemHandler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bad id value: %s", err)})
		return
	}
	s := h.GetItemService()
	if err := s.DeleteItem(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}