package main

import (
	"fmt"
	"home-manager/server/internal/application/routers"
	"home-manager/server/internal/infrastructure/db"
	"os"
)

func main() {
	test := os.Getenv("env")
	fmt.Printf("**********%s**********", test)
	db := db.Init()
	router := routers.SetupRouter(db)

	router.Run("localhost:8080")
}
