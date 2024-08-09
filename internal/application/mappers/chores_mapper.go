package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/valueobjects"
)

func ToChoreListResponse(chores []*valueobjects.Chore) []*models.Chore {
	var responseList []*models.Chore

	for _, chore := range chores {
		mapped := models.Chore{
			Id:           chore.Id(),
			Name:         chore.Name(),
			Instructions: chore.Instructions(),
			Points:       chore.Points(),
			RoomId:       chore.Points(),
		}
		responseList = append(responseList, &mapped)
	}

	return responseList
}

func ToGetAllChoresResponse(chores []valueobjects.Chore) []models.Chore {
	var responseList []models.Chore

	for _, chore := range chores {
		responseList = append(responseList, toModelChore(&chore))
	}

	return responseList
}

func ToCreateChoreResponse(chore *valueobjects.Chore) models.CreateChoreResponse {
	mapped := toModelChore(chore)

	return models.CreateChoreResponse{
		Chore: mapped,
	}
}

func ToGetChoreResponse(chore *valueobjects.Chore) models.GetChoreResponse {
	mapped := toModelChore(chore)

	return models.GetChoreResponse{
		Chore: mapped,
	}
}

func toModelChore(chore *valueobjects.Chore) models.Chore {
	return models.Chore{
		Id:           chore.Id(),
		Name:         chore.Name(),
		Instructions: chore.Instructions(),
		Points:       chore.Points(),
		RoomId:       chore.Points(),
	}
}


