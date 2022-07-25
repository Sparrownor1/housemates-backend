package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func validateToken(tokenString string) (interface{}, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	jwksUrl := fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain)
	jwks, err := keyfunc.Get(jwksUrl, keyfunc.Options{
		RefreshRateLimit: time.Second * 12,
	})
	if err != nil {
		log.Fatalf("Failed to get the JWKS from the given URL.\nError: %s", err)
	}

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return token, nil
}
