package repositories

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/core/internalerrors"
)

var userId int32 = 0

type inMemUserRepo struct {
	db []*entities.User
}

func NewInMemUserRepo() UserRepo {
	a, _ := entities.NewUser(userId, "Foo", 0)
	userId++

	users := []*entities.User{
		&a,
	}
	return &inMemUserRepo{users}
}

func (repo *inMemUserRepo) Create(user *entities.User) (entities.User, error) {
	newUser, err := entities.NewUser(id, user.Name(), user.TotalPoints())
	if err != nil {
		return entities.User{}, nil
	}

	repo.db = append(repo.db, &newUser)
	id++

	return newUser, nil
}

func (repo *inMemUserRepo) GetAll() ([]*entities.User, error) {
	return repo.db, nil
}

func (repo *inMemUserRepo) GetById(id int32) (entities.User, error) {
	for _, user := range repo.db {
		if user.Id() == id {
			return *user, nil
		}
	}

	return entities.User{}, internalerrors.NotFoundError
}