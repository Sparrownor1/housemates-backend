package auth

import (
	"crypto/rand"
	"encoding/base64"
	"housemates/housemates-backend/core/authenticator"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	auth := authenticator.GetAuthenticator()

	state, err := generateRandomState()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Save the state inside the session.
	session := sessions.Default(c)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func Register(c *gin.Context) {
	// TODO: check for password robustness
	c.String(http.StatusNotImplemented, "Register Stub")
}
