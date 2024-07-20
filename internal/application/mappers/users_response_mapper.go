package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
)

func ToUserListResponse(users []*entities.User) []*models.UserResponse {
	var responseList []*models.UserResponse

	for _, user := range users {
		mapped := models.FromUser(user)
		responseList = append(responseList, &mapped)
	}

	return responseList
}