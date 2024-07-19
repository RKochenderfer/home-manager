package chorescontroller

import (
	"home-manager/server/internal/application/mappers"
	choreservice "home-manager/server/internal/infrastructure/services/chore"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChoresController struct {
	choreService choreservice.ChoresService
}

func NewChoresController(cs choreservice.ChoresService) ChoresController {
	return ChoresController{cs}
}

func (c *ChoresController) GetChores(context *gin.Context) {
	chores, _ := c.choreService.GetAll()
	choresResponse := mappers.ToChoreListResponse(chores)

	context.IndentedJSON(http.StatusOK, choresResponse)
}
