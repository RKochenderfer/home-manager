package services

import "home-manager/server/internal/core/entities"

type ChoresService interface {
	GetAll() ([]*entities.Chore, error)
	GetById(id int32) (entities.Chore, error)
	Create(chore *entities.Chore) (entities.Chore, error)
}