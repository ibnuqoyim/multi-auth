//go:build wireinject
// +build wireinject

package main

import (
	"gihtub.com/ibnuqoyim/multi-auth/configs"
	"gihtub.com/ibnuqoyim/multi-auth/controllers"
	"gihtub.com/ibnuqoyim/multi-auth/repositories"
	"gihtub.com/ibnuqoyim/multi-auth/services"

	"github.com/google/wire"
)

func InitializeAuthController() *controllers.AuthController {
	wire.Build(
		configs.LoadConfig,
		repositories.NewUserRepository,
		services.NewAuthService,
		controllers.NewAuthController,
	)
	return &controllers.AuthController{}
}
