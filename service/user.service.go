package service

import (
	"errors"
	"github.com/prithuadhikary/user-service/domain"
	"github.com/prithuadhikary/user-service/model"
	"github.com/prithuadhikary/user-service/repository"
)

type UserService interface {
	Signup(request *model.SignupRequest) error
}

type userService struct {
	repository repository.UserRepository
}

func (service *userService) Signup(request *model.SignupRequest) error {
	if request.Password != request.PasswordConfirmation {
		return errors.New("password and confirm password must match")
	}
	exists := service.repository.ExistsByUsername(request.Username)
	if exists {
		return errors.New("email already exists")
	}
	service.repository.Save(&domain.User{
		Username: request.Username,
		Password: request.Password,
		Role:     "END_USER",
	})
	return nil
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}
