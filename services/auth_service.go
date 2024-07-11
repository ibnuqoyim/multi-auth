package services

import "gihtub.com/ibnuqoyim/multi-auth/models"

func Register(user models.User) error {
	// Simulate user registration logic
	return nil
}

func Login(email, password string) (string, error) {
	// Simulate login logic
	return "jwt_token", nil
}
