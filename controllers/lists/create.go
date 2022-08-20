package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateList(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "create list")
}
