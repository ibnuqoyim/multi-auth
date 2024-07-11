package utils

import "github.com/labstack/echo/v4"

func JSONResponse(c echo.Context, message string) error {
    return c.JSON(200, map[string]string{
        "message": message,
    })
}
