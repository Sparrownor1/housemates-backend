package middleware

import (
	"fmt"
	"housemates/housemates-backend/core/authenticator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := validateToken(c.Request.Header["Bearer"][0])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func validateToken(tokenString string) (interface{}, error) {
	jwks := authenticator.GetJWKS()

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}
