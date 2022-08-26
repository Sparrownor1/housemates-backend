// utils for auth
// hashing, jwt handling, etc
package auth

import (
	"fmt"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Password Hashing
// TODO: make this hash well
func Hash(password string) string {
	return password
}

// userClaims for jwt
type UserClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

func getSigningKey() []byte {
	godotenv.Load()

	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("no SIGNING_KEY env var")
	}

	return []byte(signingKey)
}

func ValidateTokenString(tokenString string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getSigningKey(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		id := claims.ID

		var user models.User

		result := db.GetDB().First(&user, id)

		if result.Error != nil {
			return nil, result.Error
		}

		return &user, nil
	} else {
		return nil, fmt.Errorf("not logged in: %w", err)
	}
}

func GenerateTokenString(user models.User) string {
	userClaims := UserClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
			Issuer:    "housemates",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	ss, err := token.SignedString(getSigningKey())
	if err != nil {
		log.Fatal(err)
	}
	return ss
}
