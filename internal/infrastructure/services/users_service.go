package services

import (
	"fmt"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/repositories"

	"github.com/google/uuid"
)

type usersService struct {
	repo repositories.UserRepo
}

func NewUsersService(repo repositories.UserRepo) (UsersService, error) {
	return &usersService{repo}, nil
}

func (u *usersService) Create(user *entities.User) (*entities.User, error) {
	dbUser, err := u.repo.Create(&models.User{
		Base: models.Base{ID: user.Id() },
		Name: user.Name(),
		TotalPoints: uint(user.TotalPoints()),
		Role: string(user.Role()),
	})

	if err != nil {
		return nil, err
	}

	return entities.From(dbUser.ID, dbUser.Name, int32(dbUser.TotalPoints), entities.Role(dbUser.Role))
}

func (u *usersService) GetAll() ([]*entities.User, error) {
	var users []*entities.User

	dbUsers, err := u.repo.GetAll()
	if err != nil {
		return users, err
	}

	for _, u := range dbUsers {
		ec, err := entities.From(u.ID, u.Name, int32(u.TotalPoints), entities.Role(u.Role))
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		users = append(users, ec)
	}

	return users, nil
}

func (u *usersService) GetById(id uuid.UUID) (*entities.User, error) {
	dbUser, err := u.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return entities.From(dbUser.ID, dbUser.Name, int32(dbUser.TotalPoints), entities.Role(dbUser.Role))
}
