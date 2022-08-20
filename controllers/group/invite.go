package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InviteToGroup(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "invite to group")
}
