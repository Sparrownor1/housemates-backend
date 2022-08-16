package routes

import (
	"housemates/housemates-backend/controllers/group"

	"github.com/gin-gonic/gin"
)

func addGroupRoutes(rg *gin.RouterGroup) {
	groupRoutes := rg.Group("/group")

	groupRoutes.POST("/create", group.CreateGroup)
	groupRoutes.DELETE("/delete", group.DeleteGroup)
	groupRoutes.POST("/invite", group.InviteToGroup)
	groupRoutes.POST("/join", group.JoinGroup)
}
