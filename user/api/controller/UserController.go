package controller

import (
	"ecommerce-microservice/user/api/middlewares"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/common/models"
	"ecommerce-microservice/user/usecase"
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
	group.POST("/user", u.AddUser)
	group.PATCH("/user/:id", u.UpdateUser)
	group.GET("/user", u.FetchAllUser)
	group.GET("/user/:id", u.FetchUser)
	group.DELETE("/user/:id", u.DeleteUser)
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
// @Param services body dto.RequestUser true "Login Authentication"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /user [post]
func (u *userController) AddUser(ctx echo.Context) error {
	request := dto.RequestUser{}
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

// FetchAllUser : Fetch All User
// @Summary Fetch All User
// @Description This endpoint for fetch all user
// @Tags User
// @Accept  json
// @Produce  json
// @Param services query dto.RequestPagination true "Parameter with pagination"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.User} "desc"
// @Router /user [get]
func (u *userController) FetchAllUser(ctx echo.Context) error {
	request := dto.RequestPagination{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := u.userService.FetchAllUser(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchUser : Fetch User  by ID
// @Summary Fetch User  by ID
// @Description This endpoint for fetch user  by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.User} "desc"
// @Router /user/{id} [get]
func (u *userController) FetchUser(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := u.userService.FetchUser(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateUser : Update User  by ID
// @Summary Update User  by ID
// @Description This endpoint for Update user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Param services body dto.RequestUser true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.User} "desc"
// @Router /user/{id} [PATCH]
func (u *userController) UpdateUser(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	request := dto.RequestUser{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := u.userService.UpdateUser(converted, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteUser : Delete User  by ID
// @Summary Delete User  by ID
// @Description This endpoint for Delete user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.User} "desc"
// @Router /user/{id} [DELETE]
func (u *userController) DeleteUser(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := u.userService.DeleteUser(converted)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
