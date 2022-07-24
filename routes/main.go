package routes

import (
	housemates_sessions "housemates/housemates-backend/core/sessions"

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
	api := router.Group("/api")

	// sessions
	api.Use(sessions.Sessions("housemates_session", housemates_sessions.GetStore()))

	v1 := api.Group("/v1")
	addAuthRoutes(v1)
}
