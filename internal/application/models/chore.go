package models

import "home-manager/server/internal/core/entities"

type ChoreResponse struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}

func FromChore(c *entities.Chore) ChoreResponse {
	return ChoreResponse{c.Id(), c.Name(), c.Instructions(), c.Points(), c.RoomId()}
}

type NewChoreRequest struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}
