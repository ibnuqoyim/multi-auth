package repositories

import (
	"context"

	"gihtub.com/ibnuqoyim/multi-auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	// Add other repository methods
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (repo *userRepository) CreateUser(user *models.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	return err
}
