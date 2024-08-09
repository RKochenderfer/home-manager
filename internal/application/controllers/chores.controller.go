package controllers

import (
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/valueobjects"
	"home-manager/server/internal/infrastructure/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChoresController struct {
	choreService services.ChoresService
}

func NewChoresController(cs services.ChoresService) ChoresController {
	return ChoresController{cs}
}

func (cc *ChoresController) GetAll(ctx *gin.Context) {
	chores, _ := cc.choreService.GetAll()
	getAllChoresResponse := mappers.ToGetAllChoresResponse(chores)

	ctx.JSON(http.StatusOK, getAllChoresResponse)
}

func (cc *ChoresController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return
	}

	chore, err := cc.choreService.GetById(int32(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		mapped := mappers.ToGetChoreResponse(&chore)
		ctx.JSON(http.StatusOK, mapped)
	}
}

func (cc *ChoresController) Create(ctx *gin.Context) {
	var newChoreReq models.CreateChoreRequest
	if err := ctx.BindJSON(&newChoreReq); err != nil {
		return
	}

	vo, err := valueobjects.NewChore(
		0,
		newChoreReq.Name,
		newChoreReq.Instructions,
		newChoreReq.Points,
		newChoreReq.RoomId,
	)
	if err != nil {
		return
	}
	newChore := vo.(valueobjects.Chore)
	
	chore, err := cc.choreService.Create(&newChore)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusCreated, mappers.ToCreateChoreResponse(&chore))
	}
}
