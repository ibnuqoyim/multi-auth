package main

import (
	"gihtub.com/ibnuqoyim/multi-auth/controllers"
	"gihtub.com/ibnuqoyim/multi-auth/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middlewares.AuthMiddleware)

	// Routes
	e.POST("/api/register", controllers.Register)
	e.POST("/api/login", controllers.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
