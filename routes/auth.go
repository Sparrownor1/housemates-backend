package routes

import (
	"housemates/housemates-backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

// login and register
func addAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", auth.Login)
	rg.POST("/register", auth.Register)
}
