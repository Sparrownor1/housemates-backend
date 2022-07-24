package routes

import (
	housemates_sessions "housemates/housemates-backend/core/sessions"
	"housemates/housemates-backend/middleware"

	"github.com/gin-contrib/sessions"
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
	// sessions
	router.Use(sessions.Sessions("housemates_session", housemates_sessions.GetStore()))
	// auth middleware
	router.Use(middleware.AuthMiddleware())
}
