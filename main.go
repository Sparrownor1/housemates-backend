package main

import (
	"housemates/housemates-backend/core/db"
	"housemates/housemates-backend/routes"
)

func main() {
	db.Init()
	// routes needs authenticator
	routes.Run()
}
