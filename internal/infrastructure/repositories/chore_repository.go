package repositories

import (
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/db/models"
)

type sqliteChoreRepo struct {
	db *db.Database
}

func NewSqliteChoreRepo(db *db.Database) ChoreRepo {
	return &sqliteChoreRepo{db}
}

// Create implements ChoreRepo.
func (s *sqliteChoreRepo) Create(toAdd *models.Chore) (models.Chore, error) {
	if result := s.db.Connection().Create(&toAdd); result.Error != nil {
		return models.Chore{}, nil
	}

	return s.GetById(int32(toAdd.ID))
}

// GetAll implements ChoreRepo.
func (s *sqliteChoreRepo) GetAll() ([]models.Chore, error) {
	var chores []models.Chore

	if result := s.db.Connection().Find(&chores); result.Error != nil {
		return nil, result.Error
	}

	return chores, nil
}

// GetById implements ChoreRepo.
func (s *sqliteChoreRepo) GetById(id int32) (models.Chore, error) {
	var chore models.Chore
	if result := s.db.Connection().First(&chore, id); result.Error != nil {
		return models.Chore{}, result.Error
	}

	return chore, nil
}
