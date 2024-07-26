package models

import "home-manager/server/internal/core/valueobjects"

type ChoreResponse struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}

func ChoreResponseFrom(c *valueobjects.Chore) ChoreResponse {
	return ChoreResponse{
		Id: c.Id(),
		Name: c.Name(),
		Instructions: c.Instructions(),
		Points: c.Points(),
		RoomId: c.RoomId(),
	}
}

type NewChoreRequest struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}
