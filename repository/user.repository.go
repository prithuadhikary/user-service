package repository

import (
	"github.com/prithuadhikary/user-service/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User)
	ExistsByUsername(username string) bool
}

type userRepository struct {
	db *gorm.DB
}

func (repository *userRepository) Save(user *domain.User) {
	repository.db.Save(user)
}

func (repository *userRepository) ExistsByUsername(username string) bool {
	var count int64
	repository.db.Model(&domain.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func NewUserRepository(db *gorm.DB) UserRepository {
	var repository UserRepository

	repository = &userRepository{
		db: db,
	}

	return repository
}
