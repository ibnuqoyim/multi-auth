package services

import (
	"gihtub.com/ibnuqoyim/multi-auth/models"
	"gihtub.com/ibnuqoyim/multi-auth/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(user *models.User) error
	// Add other service methods
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return service.userRepository.CreateUser(user)
}
