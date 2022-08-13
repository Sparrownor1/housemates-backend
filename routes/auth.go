package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// login and register
func addAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "login") })
	rg.POST("/register", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "register") })
}
