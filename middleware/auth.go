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
		if tokenStrings, ok := c.Request.Header["Bearer"]; ok && (len(tokenStrings) > 1) {
			token, err := validateToken(tokenStrings[0])
			if err != nil || !token.Valid {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	jwks := authenticator.GetJWKS()

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}
