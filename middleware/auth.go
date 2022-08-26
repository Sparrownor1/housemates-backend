package middleware

import (
	"housemates/housemates-backend/libs/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	TokenString string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		tokenStringHeader := strings.Split(h.TokenString, "Bearer ")

		if len(tokenStringHeader) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "must provide Authorization header with the formation `Bearer {token}`")
			return
		}

		log.Println("token:", tokenStringHeader[1])

		if user, err := auth.ValidateTokenString(tokenStringHeader[1]); err != nil {
			log.Println("auth error:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			c.Set("user", user)
		}

		c.Next()
	}
}
