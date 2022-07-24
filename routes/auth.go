package routes

import (
	authController "housemates/housemates-backend/controllers/auth"
	"housemates/housemates-backend/middleware"

	"github.com/gin-gonic/gin"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	authRouter := rg.Group("/auth")

	authRouter.GET("/login", authController.Login)
	authRouter.GET("/callback", authController.Callback)
	authRouter.GET("/logout", authController.Logout)

	// auth middleware
	rg.Use(middleware.AuthMiddleware())
}
