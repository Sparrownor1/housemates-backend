package routes

import (
	"housemates/housemates-backend/middleware"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":5001")
}

// getRoutes will create our routes of our entire application
func getRoutes() {
	api := router.Group("/api")

	v1 := api.Group("/v1")
	addAuthRoutes(v1)
	v1.Use(middleware.AuthMiddleware())
	addGroupRoutes(v1)
	addListsRoutes(v1)
}
