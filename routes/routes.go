package routes

import (
	"gihtub.com/ibnuqoyim/multi-auth/controllers"
	"go.mongodb.org/mongo-driver/mongo"

	"gihtub.com/ibnuqoyim/multi-auth/configs"
	"gihtub.com/ibnuqoyim/multi-auth/repositories"
	"gihtub.com/ibnuqoyim/multi-auth/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, config *configs.Config, db *mongo.Database) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	e.POST("/api/register", authController.Register)
	// Add other routes
}
