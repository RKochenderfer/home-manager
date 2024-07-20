package choreservice

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/repositories"
	"home-manager/server/internal/infrastructure/services"
)

type ChoresService struct {
	repo repositories.ChoreRepo
}

func NewChoresService(repo repositories.ChoreRepo) (services.ChoresService, error) {
	return &ChoresService{repo}, nil
}

func (cs *ChoresService) GetAll() ([]*entities.Chore, error) {
	return cs.repo.GetAll()
}

func (cs *ChoresService) GetById(id int32) (entities.Chore, error) {
	return cs.repo.GetById(id)
}

func (cs *ChoresService) Create(chore *entities.Chore) (entities.Chore, error) {
	return cs.repo.Create(chore)
}
