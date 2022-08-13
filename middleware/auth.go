package middleware

import (
	"fmt"
	"housemates/housemates-backend/core/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tokenStrings, ok := c.Request.Header["Bearer"]; ok && (len(tokenStrings) > 1) {
			_, err := validateTokenString(tokenStrings[0])
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// userClaims for jwt
type UserClaims struct {
	models.User
	jwt.StandardClaims
}

var signingKey = []byte("mysecret")

func validateTokenString(tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return &claims.User, nil
	} else {
		return nil, fmt.Errorf("not logged in: %w", err)
	}
}

func GenerateTokenString(user models.User) string {
	userClaims := UserClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 150000,
			Issuer:    "housemates",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		log.Fatal(err)
	}
	return ss
}
