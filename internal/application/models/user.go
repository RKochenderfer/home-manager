package models

import (
	"home-manager/server/internal/core/entities"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id          uuid.UUID  `json:"id"`
	Name        string `json:"name"`
	TotalPoints int32  `json:"totalPoints"`
	Role        string `json:"role"`
}

func FromUser(u *entities.User) UserResponse {
	return UserResponse{u.Id(), u.Name(), u.TotalPoints(), string(u.Role())}
}

type NewUserRequest struct {
	Name string        `json:"name"`
	Role entities.Role `json:"role"`
}