package main

import (
	"housemates/housemates-backend/core/authenticator"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/sessions"
	"housemates/housemates-backend/routes"
)

func main() {
	db.Init()
	sessions.Init()
	authenticator.Init()
	routes.Run()
}
