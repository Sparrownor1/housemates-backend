package authenticator

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/joho/godotenv"
)

var jwks *keyfunc.JWKS = nil

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	jwksUrl := fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain)
	jwks, err = keyfunc.Get(jwksUrl, keyfunc.Options{
		RefreshRateLimit: time.Second * 12,
	})
	if err != nil {
		log.Fatalf("Failed to get the JWKS from the given URL.\nError: %s", err)
	}
}

func GetJWKS() *keyfunc.JWKS {
	if jwks == nil {
		log.Fatal("authenticator not initialized")
	}
	return jwks
}
