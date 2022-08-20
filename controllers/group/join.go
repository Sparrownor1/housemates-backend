package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "join group")
}
