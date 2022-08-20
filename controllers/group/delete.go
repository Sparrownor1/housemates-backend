package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "delete group")
}
