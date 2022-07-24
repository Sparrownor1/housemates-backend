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
	router.Run(":5000")
}

// getRoutes will create our routes of our entire application
func getRoutes() {
	router.Use(middleware.AuthMiddleware())
}
