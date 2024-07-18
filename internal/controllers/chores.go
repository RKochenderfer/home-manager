package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"home-manager/server/internal/models"
)

var chores = []models.Chore{
	{Id: "1", Name: "Floorboards", Description: "Clean the dust off the floorboards", Points: 10},
	{Id: "2", Name: "Pickup Toys", Description: "Pick toys off floor", Points: 8},
	{Id: "3", Name: "Foo", Description: "Bar", Points: 1},
}


func GetChores(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, chores)
}