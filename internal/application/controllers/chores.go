package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/chore"
)

var a, err = chore.NewChore(1, "Floorboards", "Clean the dust off the floorboards", 10)
var b, _ = chore.NewChore(2, "Pickup Toys", "Pick toys off floor", 8)
var c, _ = chore.NewChore(3, "Foo", "Clean the dust off the floorboards", 1)

var choreResponse = models.FromChore(a)
var chores = []models.ChoreResponse{
	models.FromChore(a),
	models.FromChore(b),
	models.FromChore(c),
}

func GetChores(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, chores)
}
