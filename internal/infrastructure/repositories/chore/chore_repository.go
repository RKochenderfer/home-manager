package chorerepo

import (
	"fmt"
	"home-manager/server/internal/core/chore"
)

type ChoreRepo interface {
	GetById(id int32) (*chore.Chore, error)
	GetAll() ([]*chore.Chore, error)
}

type InMemChoreRepo struct {
	db []*chore.Chore
}

func NewInMemoChoreRepo() ChoreRepo {
	a, _ := chore.NewChore(1, "Floorboards", "Clean the dust off the floorboards", 10)
	b, _ := chore.NewChore(2, "Pickup Toys", "Pick toys off floor", 8)
	c, _ := chore.NewChore(3, "Foo", "Clean the dust off the floorboards", 1)

	chores := []*chore.Chore{
		a,
		b,
		c,
	}
	return &InMemChoreRepo{chores}
}

func (repo *InMemChoreRepo) GetById(id int32) (*chore.Chore, error) {
	for _, chore := range repo.db {
		if chore.Id() == id {
			return chore, nil
		}
	}

	return nil, fmt.Errorf("chore with id: %v was not found", id)
}

func (repo *InMemChoreRepo) GetAll() ([]*chore.Chore, error) {
	return repo.db, nil
}