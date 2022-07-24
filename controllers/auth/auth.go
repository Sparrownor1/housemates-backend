package auth

import (
	"crypto/rand"
	"encoding/base64"
	"housemates/housemates-backend/core/authenticator"
	"net/http"
	"net/url"
	"os"

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

func Callback(c *gin.Context) {
	auth := authenticator.GetAuthenticator()

	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		c.String(http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		c.String(http.StatusUnauthorized, "Failed to exchange an authorization code for a token.")
		return
	}

	idToken, err := auth.VerifyIDToken(c.Request.Context(), token)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to logged in page.
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Logout(c *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request.Host)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

func Register(c *gin.Context) {
	// TODO: check for password robustness
	c.String(http.StatusNotImplemented, "Register Stub")
}
