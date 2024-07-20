package chorerepo

import (
	"home-manager/server/internal/core/entities"

	"home-manager/server/internal/core/internalerrors"
)

type ChoreRepo interface {
	GetAll() ([]*entities.Chore, error)
	GetById(id int32) (entities.Chore, error)
}

type InMemChoreRepo struct {
	db []*entities.Chore
}

func NewInMemoChoreRepo() ChoreRepo {
	a, _ := entities.NewChore(1, "Floorboards", "Clean the dust off the floorboards", 10)
	b, _ := entities.NewChore(2, "Pickup Toys", "Pick toys off floor", 8)
	c, _ := entities.NewChore(3, "Foo", "Clean the dust off the floorboards", 1)

	chores := []*entities.Chore{
		a,
		b,
		c,
	}
	return &InMemChoreRepo{chores}
}

func (repo *InMemChoreRepo) GetAll() ([]*entities.Chore, error) {
	return repo.db, nil
}

func (repo *InMemChoreRepo) GetById(id int32) (entities.Chore, error) {
	for _, chore := range repo.db {
		if chore.Id() == id {
			return *chore, nil
		}
	}

	return entities.Chore{}, internalerrors.NotFoundError
}
