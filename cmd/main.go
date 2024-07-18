package main

import (
	"home-manager/server/internal/routers"
)

func main() {
	router := routers.SetupRouter()

	router.Run("localhost:8080")
}
