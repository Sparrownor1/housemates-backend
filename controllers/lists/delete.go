package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "delete list")
}
