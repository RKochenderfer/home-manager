package chorescontroller

import (
	"errors"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/internalerrors"
	choreservice "home-manager/server/internal/infrastructure/services/chore"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChoresController struct {
	choreService choreservice.ChoresService
}

func NewChoresController(cs choreservice.ChoresService) ChoresController {
	return ChoresController{cs}
}

func (choresController *ChoresController) GetAll(ctx *gin.Context) {
	chores, _ := choresController.choreService.GetAll()
	choresResponse := mappers.ToChoreListResponse(chores)

	ctx.IndentedJSON(http.StatusOK, choresResponse)
}

func (choresController *ChoresController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Unable to parse id to integer"))
	}

	chore, err := choresController.choreService.GetById(int32(id))

	if errors.Is(err, internalerrors.NotFoundError) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "chore was not found"})
	} else {
		mapped := models.FromChore(&chore)
		ctx.IndentedJSON(http.StatusFound, mapped)
	}
}
