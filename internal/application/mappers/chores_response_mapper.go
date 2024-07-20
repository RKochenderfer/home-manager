package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
)

func ToChoreListResponse(chores []*entities.Chore) []*models.ChoreResponse {
	var responseList []*models.ChoreResponse
	
	for _, chore := range chores {
		mapped := models.FromChore(chore)
		responseList = append(responseList, &mapped)
	}

	return responseList
}