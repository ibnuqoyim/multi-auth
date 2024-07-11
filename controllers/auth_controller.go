package controllers

import "github.com/labstack/echo/v4"

func Register(c echo.Context) error {
    return c.JSON(200, map[string]string{
        "message": "User registered successfully",
    })
}

func Login(c echo.Context) error {
    return c.JSON(200, map[string]string{
        "token": "jwt_token",
    })
}
