package controller

import (
	"ecommerce-microservice/product/api/middlewares"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/common/models"
	"ecommerce-microservice/product/usecase"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

// UserController : Controller for User Api
type userController struct {
	userService usecase.UserService
}

// NewUserController : Instance for register  UserService
func NewUserController(userService usecase.UserService) *userController {
	return &userController{userService: userService}
}

// Route : Setting User Route
func (u *userController) Route(group *echo.Group) {
	group.POST("/user/authentications", u.LoginAuth)
	group.PUT("/user/authentications", u.RefreshAuth, middleware.KeyAuth(middlewares.RefreshCheck))
	group.POST("/users", u.AddUser)
	group.DELETE("/user/authentications", u.DeleteAuth, middleware.KeyAuth(middlewares.RefreshCheck))
}

// LoginAuth godoc
// @Summary Login Authentication
// @Description This endpoint for Login Authentication
// @Tags User
// @Accept  json
// @Produce  json
// @Param services body dto.RequestLogin true "Login Authentication"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /user/authentications [post]
func (u *userController) LoginAuth(ctx echo.Context) error {
	request := dto.RequestLogin{}
	err := models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	result, err := u.userService.CheckLogin(request.Username, request.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(result.Code, result)
}

// RefreshAuth godoc
// @Summary Refresh Authentication
// @Description This endpoint for Refresh Authentication
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /user/authentications [put]
func (u *userController) RefreshAuth(ctx echo.Context) error {
	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	tempAuthID := ctx.Get("auth_id")
	authID, err := uuid.Parse(fmt.Sprint(tempAuthID))
	if err != nil {
		return err
	}

	result, err := u.userService.RefreshToken(conv, authID)
	if err != nil {
		return err
	}
	return ctx.JSON(result.Code, result)
}

// AddUser godoc
// @Summary Add User
// @Description This endpoint for Add User
// @Tags User
// @Accept  json
// @Produce  json
// @Param services body dto.RequestAddUser true "Login Authentication"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /users [post]
func (u *userController) AddUser(ctx echo.Context) error {
	request := dto.RequestAddUser{}
	err := models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	result, err := u.userService.AddUser(request)
	if err != nil {
		return err
	}
	return ctx.JSON(result.Code, result)
}

// DeleteAuth godoc
// @Summary Delete Authentication
// @Description This endpoint for Delete Authentication
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /user/authentications [delete]
func (u *userController) DeleteAuth(ctx echo.Context) error {
	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	tempAuthID := ctx.Get("auth_id")
	authID, err := uuid.Parse(fmt.Sprint(tempAuthID))
	if err != nil {
		return err
	}

	result, err := u.userService.DeleteAuth(conv, authID)
	if err != nil {
		return err
	}
	return ctx.JSON(result.Code, result)
}
