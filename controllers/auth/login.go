package auth

import (
	"errors"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/models"
	"housemates/housemates-backend/libs/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInfo struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var loginInfo LoginInfo

	if ctx.ShouldBind(&loginInfo) == nil {
		db := db.GetDB()
		email := loginInfo.Email
		// TODO: don't expect raw password
		password := loginInfo.Password
		var user models.User

		// check if user in db
		if err := db.First(&user, "email = ?", email).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// user not found
				log.Println("user not found")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				// some other error
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return
			}
		}

		// user in db
		// check if passwords match
		if auth.Hash(password) != user.PasswordHash {
			// no match
			log.Println("no password match")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.JSON(http.StatusOK, user)
		return
	}

	log.Println("login info not present")
	ctx.AbortWithStatusJSON(http.StatusBadRequest, "login info not present")
}
