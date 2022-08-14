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

type RegisterInfo struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var registerInfo RegisterInfo

	if ctx.ShouldBind(&registerInfo) == nil {
		db := db.GetDB()
		email := registerInfo.Email
		var user models.User

		// check if user in db
		if err := db.First(&user, "email = ?", email).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// user not found
				// we can create a new user
				// TODO: don't expect raw password
				user = models.User{
					FirstName:    registerInfo.FirstName,
					LastName:     registerInfo.LastName,
					Email:        registerInfo.Email,
					PasswordHash: auth.Hash(registerInfo.Password),
				}

				if err := db.Create(&user).Error; err != nil {
					// error during user creation
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
					return
				}

				// successful, log user in
				ctx.JSON(http.StatusOK, user)
				return
			} else {
				// some other error
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
				return
			}
		}

		// user in db
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "user already exists")
		return
	}

	log.Println("register info not present")
	ctx.AbortWithStatusJSON(http.StatusBadRequest, "register info not present")
}
