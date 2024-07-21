package services

import (
	"fmt"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/repositories"
)

type choresService struct {
	repo repositories.ChoreRepo
}

func NewChoresService(repo repositories.ChoreRepo) (ChoresService, error) {
	return &choresService{repo}, nil
}

func (cs *choresService) GetAll() ([]*entities.Chore, error) {
	var chores []*entities.Chore

	dbChores, err := cs.repo.GetAll()
	if err != nil {
		return chores, err
	}

	for _, c := range dbChores {
		ec, err := entities.NewChore(int32(c.ID), c.Name, c.Instructions, int32(c.Points), int32(c.RoomID))
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		chores = append(chores, &ec)
	}

	return chores, nil
}

func (cs *choresService) GetById(id int32) (entities.Chore, error) {
	dbChore, err := cs.repo.GetById(id)
	if err != nil {
		return entities.Chore{}, err
	}

	return entities.NewChore(int32(dbChore.ID), dbChore.Name, dbChore.Instructions, int32(dbChore.Points), int32(dbChore.RoomID))
}

func (cs *choresService) Create(chore *entities.Chore) (entities.Chore, error) {
	dbChore, err := cs.repo.Create(&models.Chore{
		Name: chore.Name(),
		Instructions: chore.Instructions(),
		Points: uint(chore.Points()),
		RoomID: uint(chore.RoomId()),
	})

	if err != nil {
		return entities.Chore{}, err
	}

	return entities.NewChore(int32(dbChore.ID), dbChore.Name, dbChore.Instructions, int32(dbChore.Points), int32(dbChore.RoomID))
}
