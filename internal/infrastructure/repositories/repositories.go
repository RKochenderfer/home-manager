package repositories

import "home-manager/server/internal/infrastructure/db/models"

type ChoreRepo interface {
	GetAll() ([]*models.Chore, error)
	GetById(id int32) (models.Chore, error)
	Create(chore *models.Chore) (models.Chore, error)
}

type UserRepo interface {
	GetAll() ([]*models.User, error)
	GetById(id int32) (models.User, error)
	Create(user *models.User) (models.User, error)
}