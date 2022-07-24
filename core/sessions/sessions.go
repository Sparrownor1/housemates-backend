package sessions

import (
	"housemates/housemates-backend/core/db"

	gormsessions "github.com/gin-contrib/sessions/gorm"
)

var store gormsessions.Store

func Init() {
	store = gormsessions.NewStore(db.GetDB(), true, []byte("secret"))
}

func GetStore() gormsessions.Store {
	return store
}
