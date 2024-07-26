package services

import (
	"fmt"
	"home-manager/server/internal/core/valueobjects"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/repositories"
)

type choresService struct {
	repo repositories.ChoreRepo
}

func NewChoresService(repo repositories.ChoreRepo) (ChoresService, error) {
	return &choresService{repo}, nil
}

func (cs *choresService) GetAll() ([]*valueobjects.Chore, error) {
	var chores []*valueobjects.Chore

	dbChores, err := cs.repo.GetAll()
	if err != nil {
		return chores, err
	}

	for _, c := range dbChores {
		vo, err := valueobjects.NewChore(int32(c.ID), c.Name, c.Instructions, int32(c.Points), int32(c.RoomID))

		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		ec, _ := vo.(valueobjects.Chore)

		chores = append(chores, &ec)
	}

	return chores, nil
}

func (cs *choresService) GetById(id int32) (valueobjects.Chore, error) {
	dbChore, err := cs.repo.GetById(id)
	if err != nil {
		return valueobjects.Chore{}, err
	}

	vo, err := valueobjects.NewChore(int32(dbChore.ID), dbChore.Name, dbChore.Instructions, int32(dbChore.Points), int32(dbChore.RoomID))
	if err != nil {
		return valueobjects.Chore{}, nil
	}
	chore, _ := vo.(valueobjects.Chore)
	
	return chore, nil
}

func (cs *choresService) Create(chore *valueobjects.Chore) (valueobjects.Chore, error) {
	dbChore, err := cs.repo.Create(&models.Chore{
		Name: chore.Name(),
		Instructions: chore.Instructions(),
		Points: uint(chore.Points()),
		RoomID: uint(chore.RoomId()),
	})

	if err != nil {
		return valueobjects.Chore{}, err
	}

	vo, err := valueobjects.NewChore(int32(dbChore.ID), dbChore.Name, dbChore.Instructions, int32(dbChore.Points), int32(dbChore.RoomID))
	if err != nil {
		return valueobjects.Chore{}, nil
	}
	c, _ := vo.(valueobjects.Chore)

	return c, nil
}
