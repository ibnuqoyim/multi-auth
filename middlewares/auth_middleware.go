package middlewares

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte("secret"),
    })(next)
}
