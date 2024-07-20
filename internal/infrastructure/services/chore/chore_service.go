package choreservice

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/repositories/chore"
)

type ChoresService interface {
	GetAll() ([]*entities.Chore, error)
	GetById(id int32) (entities.Chore, error)
}

type choresService struct {
	repo chorerepo.ChoreRepo
}

func NewChoresService(repo chorerepo.ChoreRepo) (ChoresService, error) {
	return &choresService{repo}, nil
}

func (cs *choresService) GetAll() ([]*entities.Chore, error) {
	return cs.repo.GetAll()
}

func (cs *choresService) GetById(id int32) (entities.Chore, error) {
	return cs.repo.GetById(id)
}
