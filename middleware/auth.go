package middleware

import (
	"housemates/housemates-backend/libs/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tokenStrings, ok := c.Request.Header["Bearer"]; ok && (len(tokenStrings) > 1) {
			_, err := auth.ValidateTokenString(tokenStrings[0])
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
