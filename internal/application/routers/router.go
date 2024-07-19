package routers

import (
	chorescontroller "home-manager/server/internal/application/controllers"
	chorerepo "home-manager/server/internal/infrastructure/repositories/chore"
	choreservice "home-manager/server/internal/infrastructure/services/chore"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	repo := chorerepo.NewInMemoChoreRepo()
	cs, _ := choreservice.NewChoresService(repo)
	cc := chorescontroller.NewChoresController(cs)

	// Public routes
	public := router.Group("api")
	{
		public.GET("/chores", cc.GetAll)
		public.GET("/chores/:id", cc.GetById)
	}

	// Protected Routes

	return router
}
