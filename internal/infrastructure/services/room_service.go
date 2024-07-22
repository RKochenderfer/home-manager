package services

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/repositories"
)

type roomService struct {
	repo repositories.RoomRepo
}

func NewRoomService(repo repositories.RoomRepo) (RoomService, error) {
	return &roomService{repo}, nil
}

// Create implements RoomService.
func (r *roomService) Create(user *entities.Room) (entities.Room, error) {
	panic("unimplemented")
}

// GetAll implements RoomService.
func (r *roomService) GetAll() ([]*entities.Room, error) {
	panic("unimplemented")
}

// GetById implements RoomService.
func (r *roomService) GetById(id int32) (entities.Room, error) {
	panic("unimplemented")
}
