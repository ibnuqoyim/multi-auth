package main

import (
    "github.com/labstack/echo/v4"
    "login-service/configs"
    "login-service/controllers"
    "login-service/middlewares"
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
