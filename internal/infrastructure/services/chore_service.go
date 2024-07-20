package services

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/repositories"
)

type choresService struct {
	repo repositories.ChoreRepo
}

func NewChoresService(repo repositories.ChoreRepo) (ChoresService, error) {
	return &choresService{repo}, nil
}

func (cs *choresService) GetAll() ([]*entities.Chore, error) {
	return cs.repo.GetAll()
}

func (cs *choresService) GetById(id int32) (entities.Chore, error) {
	return cs.repo.GetById(id)
}

func (cs *choresService) Create(chore *entities.Chore) (entities.Chore, error) {
	return cs.repo.Create(chore)
}
