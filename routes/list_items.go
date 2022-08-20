package routes

import (
	listitems "housemates/housemates-backend/controllers/list_items"

	"github.com/gin-gonic/gin"
)

func addListItemRoutes(rg *gin.RouterGroup) {
	listItemRoutes := rg.Group("/:id")

	listItemRoutes.POST("/add", listitems.Add)
	listItemRoutes.PATCH("/update", listitems.Update)
	listItemRoutes.DELETE("/delete", listitems.Delete)
}
