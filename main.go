package main

import (
	"housemates/housemates-backend/core/authenticator"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/core/sessions"
	"housemates/housemates-backend/routes"
)

func main() {
	db.Init()
	// sessions needs db
	sessions.Init()
	// authenticator needs sessions
	authenticator.Init()
	// routes needs authenticator
	routes.Run()
}
