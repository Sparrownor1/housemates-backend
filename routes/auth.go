package routes

import (
	"housemates/housemates-backend/controllers/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login and register
func addAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", auth.Login)
	rg.POST("/register", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "register") })
}
