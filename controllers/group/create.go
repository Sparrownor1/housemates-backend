package group

import (
	"fmt"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/models"
	"housemates/housemates-backend/libs/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type groupInfo struct {
	GroupName string `binding:"required" form:"group_name" json:"group_name"`
}

func CreateGroup(ctx *gin.Context) {
	value, exists := ctx.Get("user")
	// ensure user exists
	if !exists {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("user doesn't exist after middleware"))
		return
	}
	user := value.(*models.User)

	// ensure user not already in a group
	if user.GroupID != nil {
		ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("user has group already"))
		return
	}

	// get group name from body
	var groupInfo groupInfo
	if ctx.ShouldBind(&groupInfo) != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("no field group_name supplied"))
		return
	}

	group := models.Group{
		Name:        groupInfo.GroupName,
		InviteCode:  group.GenerateInviteCode(),
		AdminUserID: user.ID,
	}

	db := db.GetDB()
	result := db.Create(&group)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error creating group record in database"))
		return
	}

	// add user to group
	err := db.Model(user).Association("Group").Replace(&group)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error adding user to group in database: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"group": group})
}
