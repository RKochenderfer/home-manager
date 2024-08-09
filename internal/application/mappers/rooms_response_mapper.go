package mappers

import (
	"home-manager/server/internal/application/models"
	"home-manager/server/internal/core/entities"
)

func ToRoomsListResponse(rooms []*entities.Room) []*models.RoomResponse {
	var responseList []*models.RoomResponse

	for _, room := range rooms {
		mapped := models.RoomResponse{Id: room.GetId(), Name: room.GetName()}
		responseList = append(responseList, &mapped)
	}

	return responseList
}
