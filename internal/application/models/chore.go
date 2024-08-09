package models

import "home-manager/server/internal/core/valueobjects"

type Chore struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}

func ChoreResponseFrom(c *valueobjects.Chore) Chore {
	return Chore{
		Id:           c.Id(),
		Name:         c.Name(),
		Instructions: c.Instructions(),
		Points:       c.Points(),
		RoomId:       c.RoomId(),
	}
}

type CreateChoreRequest struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Points       int32  `json:"points"`
	RoomId       int32  `json:"roomId"`
}

type CreateChoreResponse struct {
	Chore
}

type GetChoreResponse struct {
	Chore
}

type GetAllChoresResponse struct {
	chores []Chore
}
