package services

import (
	"fmt"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/repositories"
)

type usersService struct {
	repo repositories.UserRepo
}

func NewUsersService(repo repositories.UserRepo) (UsersService, error) {
	return &usersService{repo}, nil
}

func (u *usersService) Create(user *entities.User) (entities.User, error) {
	dbUser, err := u.repo.Create(&models.User{
		Name: user.Name(),
		TotalPoints: uint(user.TotalPoints()),
	})

	if err != nil {
		return entities.User{}, err
	}

	return entities.NewUser(int32(dbUser.ID), dbUser.Name, int32(dbUser.TotalPoints), entities.Role(dbUser.Name))
}

func (u *usersService) GetAll() ([]*entities.User, error) {
	var users []*entities.User

	dbUsers, err := u.repo.GetAll()
	if err != nil {
		return users, err
	}

	for _, u := range dbUsers {
		ec, err := entities.NewUser(int32(u.ID), u.Name, int32(u.TotalPoints), entities.Role(u.Role))
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		users = append(users, &ec)
	}

	return users, nil
}

func (u *usersService) GetById(id int32) (entities.User, error) {
	dbUser, err := u.repo.GetById(id)
	if err != nil {
		return entities.User{}, err
	}

	return entities.NewUser(int32(dbUser.ID), dbUser.Name, int32(dbUser.TotalPoints), entities.Role(dbUser.Role))
}
