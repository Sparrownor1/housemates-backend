package sessions

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions/cookie"
)

var store cookie.Store

func Init() {
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	store = cookie.NewStore([]byte("secret"))
}

func GetStore() cookie.Store {
	return store
}
