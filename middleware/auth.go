package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("profile") == nil {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			c.Next()
		}
	}
}
