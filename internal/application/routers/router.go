package routers

import (
	controllers "home-manager/server/internal/application/controllers"
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/repositories"
	"home-manager/server/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db db.Database) *gin.Engine {
	router := gin.Default()

	// Public routes
	setupChoresRoutes(router, &db)
	setupUsersRoutes(router, &db)

	// Protected Routes

	return router
}

func setupChoresRoutes(router *gin.Engine, db *db.Database) {
	repo := repositories.NewSqliteChoreRepo(db)
	cs, _ := services.NewChoresService(repo)

	cc := controllers.NewChoresController(cs)

	chores := router.Group("api/chores")
	{
		chores.GET("/", cc.GetAll)
		chores.GET("/:id", cc.GetById)
		chores.POST("/", cc.Create)
	}
}

func setupUsersRoutes(router *gin.Engine, db *db.Database) {
	repo := repositories.NewSqliteUserRepo(db)
	us, _ := services.NewUsersService(repo)

	uc := controllers.NewUsersController(us)

	chores := router.Group("api/users")
	{
		chores.GET("/", uc.GetAll)
		chores.GET("/:id", uc.GetById)
		chores.POST("/", uc.Create)
	}
}
