package authenticator

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/jwt"
)

var jwtConfig *jwt.Config

func Init() {
	jwtConfig = new(jwt.Config)

	var err error

	// get dbinfo from .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	email := os.Getenv("AUTH0_EMAIL")
}
