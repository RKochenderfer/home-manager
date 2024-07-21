package controllers

import (
	"errors"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/core/internalerrors"
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
	choresResponse := mappers.ToChoreListResponse(chores)

	ctx.JSON(http.StatusOK, choresResponse)
}

func (cc *ChoresController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("unable to parse id to integer"))
	}

	chore, err := cc.choreService.GetById(int32(id))

	if errors.Is(err, internalerrors.NotFoundError) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "chore was not found"})
	} else {
		mapped := models.FromChore(&chore)
		ctx.JSON(http.StatusOK, mapped)
	}
}

func (cc *ChoresController) Create(ctx *gin.Context) {
	var newChoreReq models.NewChoreRequest

	if err := ctx.BindJSON(&newChoreReq); err != nil {
		return
	}

	newChore, err := entities.NewChore(0, newChoreReq.Name, newChoreReq.Instructions, newChoreReq.Points, newChoreReq.RoomId)
	if err != nil {
		return
	}

	chore, err := cc.choreService.Create(&newChore)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, models.FromChore(&chore))
}
