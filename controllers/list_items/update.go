package listitems

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "update list item")
}
