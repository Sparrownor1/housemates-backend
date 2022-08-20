package listitems

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "add list item")
}
