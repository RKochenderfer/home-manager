package controllers

import (
	"fmt"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomsController struct {
	roomService services.RoomService
}

func NewRoomsController(rs services.RoomService) RoomsController {
	return RoomsController{rs}
}

func (rc *RoomsController) GetAll(ctx *gin.Context) {
	rooms, err := rc.roomService.GetAll()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	roomsResponse := mappers.ToRoomsListResponse(rooms)

	ctx.JSON(http.StatusOK, roomsResponse)
}

func (rc *RoomsController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return
	}

	room, err := rc.roomService.GetById(int32(id))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	} else {
		mapped := models.RoomResponse{Id: room.GetId(), Name: room.GetName()}
		ctx.JSON(http.StatusOK, mapped)
	}
}

func (rc *RoomsController) Create(ctx *gin.Context) {
	var newRoomReq models.RoomRequest

	if err := ctx.BindJSON(&newRoomReq); err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	newRoom, err := entities.NewRoom(0, newRoomReq.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	room, err := rc.roomService.Create(&newRoom)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &room)
		return
	}
	ctx.JSON(http.StatusCreated, models.RoomResponse{Id: room.GetId(), Name: room.GetName()})
}

func (rc *RoomsController) CreateChore(ctx *gin.Context) {
	panic("unimplemented")
}
