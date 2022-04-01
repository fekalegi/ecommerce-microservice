package controller

import (
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/common/models"
	"ecommerce-microservice/user/usecase"
	"github.com/labstack/echo/v4"
	"strconv"
)

// RoleController : Controller for Role Api
type roleController struct {
	roleService usecase.RoleService
}

// NewRoleController : Instance for register  RoleService
func NewRoleController(roleService usecase.RoleService) *roleController {
	return &roleController{roleService: roleService}
}

// Route : Setting Role Route
func (n *roleController) Route(group *echo.Group) {
	group.POST("/role", n.CreateRole)
	group.GET("/role", n.FetchAllRole)
	group.GET("/role/:id", n.FetchRole)
	group.PATCH("/role/:id", n.UpdateRole)
	group.DELETE("/role/:id", n.DeleteRole)
}

// CreateRole godoc
// @Summary Create New Role
// @Description This endpoint for create new Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param services body dto.RequestRole true "Create new role"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /role [post]
func (n *roleController) CreateRole(ctx echo.Context) error {

	request := dto.RequestRole{}

	err := models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := n.roleService.CreateRole(request, 0)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllRole : Fetch All Role
// @Summary Fetch All Role
// @Description This endpoint for fetch all product category
// @Tags Role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Role} "desc"
// @Router /roles [get]
func (n *roleController) FetchAllRole(ctx echo.Context) error {

	responses, err := n.roleService.FetchAllRole()
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchRole : Fetch Role by ID
// @Summary Fetch Role by ID
// @Description This endpoint for fetch product category by ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Role} "desc"
// @Router /role/{id} [get]
func (n *roleController) FetchRole(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := n.roleService.FetchRole(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateRole : Update Role by ID
// @Summary Update Role by ID
// @Description This endpoint for Update product category by ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Param services body dto.RequestRole true "Create new role"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Role} "desc"
// @Router /role/{id} [put]
func (n *roleController) UpdateRole(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}
	request := dto.RequestRole{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := n.roleService.UpdateRole(converted, 0, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteRole : Delete Role by ID
// @Summary Delete Role by ID
// @Description This endpoint for Delete product category by ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Role} "desc"
// @Router /role/{id} [delete]
func (n *roleController) DeleteRole(ctx echo.Context) error {
	parameter := ctx.Param("id")
	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := n.roleService.DeleteRole(converted, 0)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
