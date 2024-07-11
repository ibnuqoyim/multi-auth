package controllers

import (
	"net/http"

	"gihtub.com/ibnuqoyim/multi-auth/models"
	"gihtub.com/ibnuqoyim/multi-auth/services"
	"gihtub.com/ibnuqoyim/multi-auth/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Data"
// @Success 200 {object} utils.Response
// @Router /register [post]
func (controller *AuthController) Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
	}

	if err := controller.authService.RegisterUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("User registered successfully"))
}

// Add other authentication methods similarly
