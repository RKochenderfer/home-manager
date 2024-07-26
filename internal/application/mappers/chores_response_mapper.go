package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/valueobjects"
)

func ToChoreListResponse(chores []*valueobjects.Chore) []*models.ChoreResponse {
	var responseList []*models.ChoreResponse
	
	for _, chore := range chores {
		mapped := models.ChoreResponse{
			Id: chore.Id(),
			Name: chore.Name(),
			Instructions: chore.Instructions(),
			Points: chore.Points(),
			RoomId: chore.Points(),
		}
		responseList = append(responseList, &mapped)
	}

	return responseList
}