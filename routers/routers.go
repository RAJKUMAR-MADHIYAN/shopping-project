package routers

// TODO: move usage of gorm

import (
	"shops/handlers"
	"shops/repository"
	"shops/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup is a function to initiate gin and define routes, it is used in test
func Setup() *gin.Engine {
	router := gin.Default()
	var r = *repository.NewItemRepository(sqlite.Open("items.db"), &gorm.Config{})
	var s = services.NewItemService(r)
	var h = handlers.NewItemHandler(*s)
	// Routes
	router.GET("/items", h.ListItems)
	router.POST("/items", h.CreateItem)
	router.PATCH("/items/:id", h.UpdateItem)
	router.GET("/items/:id", h.FindItem)
	router.DELETE("/items/:id", h.DeleteItem)
	return router
}
