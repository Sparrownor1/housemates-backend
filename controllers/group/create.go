package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create group")
}
