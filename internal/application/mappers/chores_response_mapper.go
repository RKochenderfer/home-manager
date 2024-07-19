package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/chore"
)

func ToChoreListResponse(chores []*chore.Chore) []*models.ChoreResponse {
	var responseList []*models.ChoreResponse
	
	for _, chore := range chores {
		mapped := models.FromChore(chore)
		responseList = append(responseList, &mapped)
	}

	return responseList
}