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

type UsersController struct {
	usersService services.UsersService
}

func NewUsersController(us services.UsersService) UsersController {
	return UsersController{us}
}

func (uc *UsersController) GetAll(ctx *gin.Context) {
	users, _ := uc.usersService.GetAll()
	usersResponse := mappers.ToUserListResponse(users)

	ctx.JSON(http.StatusOK, usersResponse)
}

func (uc *UsersController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("unable to parse id to integer"))
	}

	user, err := uc.usersService.GetById(int32(id))

	if errors.Is(err, internalerrors.NotFoundError) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "chore was not found"})
	} else {
		mapped := models.FromUser(&user)
		ctx.JSON(http.StatusOK, mapped)
	}
}

func (uc *UsersController) Create(ctx *gin.Context) {
	var newUserReq models.NewUserRequest

	if err := ctx.BindJSON(&newUserReq); err != nil {
		return
	}

	newUser, err := entities.NewUser(0, newUserReq.Name, 0)
	if err != nil {
		return
	}

	user, err := uc.usersService.Create(&newUser)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, models.FromUser(&user))
}