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

// ProductController : Controller for Product Api
type productController struct {
	productService usecase.ProductService
}

// NewProductController : Instance for register  ProductService
func NewProductController(productService usecase.ProductService) *productController {
	return &productController{productService: productService}
}

// Route : Setting Product Route
func (c *productController) Route(group *echo.Group) {
	group.POST("/product", c.CreateProduct, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/product", c.FetchAllProduct)
	group.GET("/product/:id", c.FetchProduct)
	group.PATCH("/product/:id", c.UpdateProduct, middleware.KeyAuth(middlewares.AuthCheck))
	group.DELETE("/product/:id", c.DeleteProduct, middleware.KeyAuth(middlewares.AuthCheck))
}

// CreateProduct godoc
// @Summary Create New Product
// @Description This endpoint for create new Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param services body dto.RequestProduct true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /product [post]
func (c *productController) CreateProduct(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestProduct{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.productService.CreateProductCommand(request, *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllProduct : Fetch All Product
// @Summary Fetch All Product
// @Description This endpoint for fetch all product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param services query dto.RequestPagination true "Parameter with pagination"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Product} "desc"
// @Router /product [get]
func (c *productController) FetchAllProduct(ctx echo.Context) error {
	request := dto.RequestPagination{}

	err := models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.productService.FetchAllProductQuery(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchProduct : Fetch Product by ID
// @Summary Fetch Product by ID
// @Description This endpoint for fetch product by ID
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Product} "desc"
// @Router /product/{id} [get]
func (c *productController) FetchProduct(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.productService.FetchProductQuery(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateProduct : Update Product by ID
// @Summary Update Product by ID
// @Description This endpoint for Update product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Param services body dto.RequestProduct true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Product} "desc"
// @Router /product/{id} [PATCH]
func (c *productController) UpdateProduct(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	request := dto.RequestUpdateProduct{}

	err = models.ValidateRequest(ctx, &request)

	if err != nil {
		return err
	}

	responses, err := c.productService.UpdateProductCommand(converted, request, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteProduct : Delete Product by ID
// @Summary Delete Product by ID
// @Description This endpoint for Delete product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Product} "desc"
// @Router /product/{id} [DELETE]
func (c *productController) DeleteProduct(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.productService.DeleteProductCommand(converted, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
