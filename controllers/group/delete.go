package group

import (
	"fmt"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteGroup(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	// ensure user exists
	if !exists {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user doesn't exist after middleware"))
		return
	}
	user := value.(*models.User)

	// ensure user already in a group
	if user.GroupID == nil {
		ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("user has no group"))
		return
	}

	db := db.GetDB()

	var group models.Group
	db.Model(user).Association("Group").Find(&group)

	if group.AdminUserID != user.ID {
		ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("user is not the admin of their group"))
		return
	}

	// this line is here as for some reason the OnDelete constraint is not being added properly
	db.Model(user).Association("Group").Delete(group)

	result := db.Delete(&group)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error deleting group from database: %s", result.Error))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}
