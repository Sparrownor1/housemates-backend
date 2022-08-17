package routes

import (
	"github.com/gin-gonic/gin"

	"housemates/housemates-backend/controllers/lists"
)

func addListsRoutes(rg *gin.RouterGroup) {
	listsRoutes := rg.Group("/lists")

	listsRoutes.GET("/", lists.Get)
	listsRoutes.GET("/all", lists.GetAll)
	listsRoutes.POST("/create", lists.CreateList)
	listsRoutes.PATCH("/update", lists.Update)
	listsRoutes.DELETE("/delete", lists.Delete)
}
