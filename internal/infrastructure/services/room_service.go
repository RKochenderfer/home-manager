package services

import (
	"fmt"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/repositories"
)

type roomService struct {
	repo repositories.RoomRepo
}

func NewRoomService(repo repositories.RoomRepo) (RoomService, error) {
	return &roomService{repo}, nil
}

// Create implements RoomService.
func (r *roomService) Create(room *entities.Room) (*entities.Room, error) {
	dbRoom := &models.Room{Name: room.GetName()}
	res, err := r.repo.Create(dbRoom)
	if err != nil {
		return nil, err
	}

	return entities.NewRoom(int32(res.ID), res.Name)
}

// GetAll implements RoomService.
func (r *roomService) GetAll() ([]*entities.Room, error) {
	var entityRooms []*entities.Room

	dbRooms, err := r.repo.GetAll()
	if err != nil {
		return entityRooms, err
	}

	for _, r := range dbRooms {
		er, err := entities.NewRoom(int32(r.ID), r.Name)
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		entityRooms = append(entityRooms, er)
	}

	return entityRooms, nil
}

// GetById implements RoomService.
func (r *roomService) GetById(id int32) (*entities.Room, error) {
	dbRoom, err := r.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return entities.NewRoom(int32(dbRoom.ID), dbRoom.Name)
}
