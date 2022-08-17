package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "get list")
}
