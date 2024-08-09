package main

import (
	"home-manager/server/internal/application/routers"
	"home-manager/server/internal/infrastructure/db"
)

func main() {
	db := db.Init()
	router := routers.SetupRouter(db)

	router.Run("localhost:8080")
}
