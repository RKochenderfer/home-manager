package services

import (
	"fmt"
	"home-manager/server/internal/core/entities/useraggregate"
	"home-manager/server/internal/infrastructure/db/models"
	"home-manager/server/internal/infrastructure/mappers"
	"home-manager/server/internal/infrastructure/repositories"

	"github.com/google/uuid"
)

type usersService struct {
	repo repositories.UserRepo
}

func NewUsersService(repo repositories.UserRepo) (UsersService, error) {
	return &usersService{repo}, nil
}

func (u *usersService) Create(user *useraggregate.User) (*useraggregate.User, error) {
	dbUser, err := u.repo.Create(&models.User{
		Base:        models.Base{ID: user.Id()},
		Name:        user.Name(),
		TotalPoints: uint(user.TotalPoints()),
		Role:        string(user.Role()),
	})

	if err != nil {
		return nil, err
	}
	ea := mappers.MapModelsAssignmentToEntityAssignments(dbUser.Assignments)

	return useraggregate.UserFrom(dbUser.ID, dbUser.Name, int32(dbUser.TotalPoints), useraggregate.Role(dbUser.Role), ea)
}

func (u *usersService) GetAll() ([]useraggregate.User, error) {
	var users []useraggregate.User

	dbUsers, err := u.repo.GetAll()
	if err != nil {
		return users, err
	}

	for _, u := range dbUsers {
		assignments := mappers.MapModelsAssignmentToEntityAssignments(u.Assignments)
		eu, err := useraggregate.UserFrom(u.ID, u.Name, int32(u.TotalPoints), useraggregate.Role(u.Role), assignments)
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		users = append(users, *eu)
	}

	return users, nil
}

func (u *usersService) GetById(id uuid.UUID) (*useraggregate.User, error) {
	dbUser, err := u.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	assignments := mappers.MapModelsAssignmentToEntityAssignments(dbUser.Assignments)
	return useraggregate.UserFrom(dbUser.ID, dbUser.Name, int32(dbUser.TotalPoints), useraggregate.Role(dbUser.Role), assignments)
}
