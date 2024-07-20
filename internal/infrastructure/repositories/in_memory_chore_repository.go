package repositories

import (
	"home-manager/server/internal/core/entities"

	"home-manager/server/internal/core/internalerrors"
)

var id int32 = 0

type inMemChoreRepo struct {
	db []*entities.Chore
}

func NewInMemChoreRepo() ChoreRepo {
	a, _ := entities.NewChore(id, "Floorboards", "Clean the dust off the floorboards", 10)
	id++
	b, _ := entities.NewChore(id, "Pickup Toys", "Pick toys off floor", 8)
	id++
	c, _ := entities.NewChore(id, "Foo", "Clean the dust off the floorboards", 1)
	id++

	chores := []*entities.Chore{
		&a,
		&b,
		&c,
	}
	return &inMemChoreRepo{chores}
}

func (repo *inMemChoreRepo) GetAll() ([]*entities.Chore, error) {
	return repo.db, nil
}

func (repo *inMemChoreRepo) GetById(id int32) (entities.Chore, error) {
	for _, chore := range repo.db {
		if chore.Id() == id {
			return *chore, nil
		}
	}

	return entities.Chore{}, internalerrors.NotFoundError
}

func (repo *inMemChoreRepo) Create(chore *entities.Chore) (entities.Chore, error) {
	newChore, err := entities.NewChore(id, chore.Name(), chore.Description(), chore.Points())
	if err != nil {
		return entities.Chore{}, err
	}

	repo.db = append(repo.db, &newChore)
	id++

	return newChore, nil
}
