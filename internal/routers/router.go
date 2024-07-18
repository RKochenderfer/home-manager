package routers

import (
	"home-manager/server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("api")
	{
		public.GET("/chores", controllers.GetChores)
	}

	// Protected Routes

	return router
}