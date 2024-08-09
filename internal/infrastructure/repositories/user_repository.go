package repositories

import (
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type sqliteUserRepo struct {
	db *db.Database
}

func NewSqliteUserRepo(db *db.Database) UserRepo {
	return &sqliteUserRepo{db}
}

// Create implements UserRepo.
func (s *sqliteUserRepo) Create(user *models.User) (models.User, error) {
	if result := s.db.Connection().Create(&user); result.Error != nil {
		return models.User{}, nil
	}

	return s.GetById(user.ID)
}

// GetAll implements UserRepo.
func (s *sqliteUserRepo) GetAll() ([]models.User, error) {
	var users []models.User

	if result := s.db.Connection().Model(&models.User{}).Preload("Assignments").Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// GetById implements UserRepo.
func (s *sqliteUserRepo) GetById(id uuid.UUID) (models.User, error) {
	var user models.User
	if result := s.db.Connection().First(&user, id); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
