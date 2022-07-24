package main

import (
	"housemates/housemates-backend/core/authenticator"
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/routes"
)

func main() {
	db.Init()
	authenticator.Init()
	routes.Run()
}
