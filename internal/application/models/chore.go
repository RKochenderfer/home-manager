package models

import "home-manager/server/internal/core/chore"

type ChoreResponse struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int32  `json:"points"`
}

func FromChore(c *chore.Chore) ChoreResponse {
	return ChoreResponse{c.Id(), c.Name(), c.Description(), c.Points()}
}

type NewChoreRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}