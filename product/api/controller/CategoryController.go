package controller

import (
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/common/helper"
	"ecommerce-microservice/product/common/models"
	"ecommerce-microservice/product/usecase"
	"github.com/fekalegi/custom-package/authentications/middlewares"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CategoryController : Controller for Category Api
type categoryController struct {
	categoryService usecase.CategoryService
}

// NewCategoryController : Instance for register  CategoryService
func NewCategoryController(categoryService usecase.CategoryService) *categoryController {
	return &categoryController{categoryService: categoryService}
}

// Route : Setting Category Route
func (c *categoryController) Route(group *echo.Group) {
	group.POST("/category", c.CreateCategory, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/category", c.FetchAllCategory)
	group.GET("/category/:id", c.FetchCategory)
	group.PATCH("/category/:id", c.UpdateCategory, middleware.KeyAuth(middlewares.AuthCheck))
	group.DELETE("/category/:id", c.DeleteCategory, middleware.KeyAuth(middlewares.AuthCheck))
}

// CreateCategory godoc
// @Summary Create New Category
// @Description This endpoint for create new Category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param services body dto.RequestCategory true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /category [post]
func (c *categoryController) CreateCategory(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestCategory{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.categoryService.CreateCategoryCommand(request, *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllCategory : Fetch All Category
// @Summary Fetch All Category
// @Description This endpoint for fetch all category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param services query dto.RequestPagination true "Parameter with pagination"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Category} "desc"
// @Router /category [get]
func (c *categoryController) FetchAllCategory(ctx echo.Context) error {
	request := dto.RequestPagination{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.categoryService.FetchAllCategoryQuery(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchCategory : Fetch Category by ID
// @Summary Fetch Category by ID
// @Description This endpoint for fetch category by ID
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id path string true "Category Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Category} "desc"
// @Router /category/{id} [get]
func (c *categoryController) FetchCategory(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.categoryService.FetchCategoryQuery(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateCategory : Update Category by ID
// @Summary Update Category by ID
// @Description This endpoint for Update category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id path string true "Category Id"
// @Param services body dto.RequestCategory true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Category} "desc"
// @Router /category/{id} [PATCH]
func (c *categoryController) UpdateCategory(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	request := dto.RequestUpdateCategory{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := c.categoryService.UpdateCategoryCommand(converted, request, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteCategory : Delete Category by ID
// @Summary Delete Category by ID
// @Description This endpoint for Delete category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id path string true "Category Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Category} "desc"
// @Router /category/{id} [DELETE]
func (c *categoryController) DeleteCategory(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.categoryService.DeleteCategoryCommand(converted, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
