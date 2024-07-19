package main

import (
	"home-manager/server/internal/application/routers"
)

func main() {
	router := routers.SetupRouter()

	router.Run("localhost:8080")
}
