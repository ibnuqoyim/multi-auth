package main

import (
	"log"
	"os"

	"gihtub.com/ibnuqoyim/multi-auth/configs"
	"gihtub.com/ibnuqoyim/multi-auth/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Login Service API
// @version 1.0
// @description This is a login service API.
// @host localhost:1323
// @BasePath /
func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	db, err := configs.ConnectDB(config.MongoURI)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e, &config, db)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	log.Fatal(e.Start(":" + port))
}
