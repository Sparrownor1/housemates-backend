package lists

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "get all lists")
}
