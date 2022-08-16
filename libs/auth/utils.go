// utils for auth
// hashing, jwt handling, etc
package auth

import (
	"fmt"
	"housemates/housemates-backend/core/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Password Hashing
// TODO: make this hash well
func Hash(password string) string {
	return password
}

// userClaims for jwt
type UserClaims struct {
	models.User
	jwt.RegisteredClaims
}

// JWT
var signingKey = []byte("mysecret")

func ValidateTokenString(tokenString string) (*models.User, error) {
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
		// TODO: cleanup what we are sending back here
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
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
