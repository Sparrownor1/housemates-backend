package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Login Stub")
}

func Register(c *gin.Context) {
	// TODO: check for password robustness
	c.String(http.StatusNotImplemented, "Register Stub")
}
