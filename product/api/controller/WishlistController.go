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

// WishlistController : Controller for Wishlist Api
type wishlistController struct {
	wishlistService usecase.WishlistService
}

// NewWishlistController : Instance for register  WishlistService
func NewWishlistController(wishlistService usecase.WishlistService) *wishlistController {
	return &wishlistController{wishlistService: wishlistService}
}

// Route : Setting Wishlist Route
func (c *wishlistController) Route(group *echo.Group) {
	group.POST("/wishlist", c.CreateWishlist, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/wishlist", c.FetchAllWishlist, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/wishlist/:id", c.FetchWishlist)
	group.DELETE("/wishlist/:id", c.DeleteWishlist, middleware.KeyAuth(middlewares.AuthCheck))
}

// CreateWishlist godoc
// @Summary Create New Wishlist
// @Description This endpoint for create new Wishlist
// @Tags Wishlist
// @Accept  json
// @Produce  json
// @Param services body dto.RequestWishlist true "Create new services"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /wishlist [post]
func (c *wishlistController) CreateWishlist(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestWishlist{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.wishlistService.CreateWishlistCommand(request, *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllWishlist : Fetch All Wishlist
// @Summary Fetch All Wishlist
// @Description This endpoint for fetch all wishlist
// @Tags Wishlist
// @Accept  json
// @Produce  json
// @Param services query dto.RequestPagination true "Parameter with pagination"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Wishlist} "desc"
// @Router /wishlist [get]
func (c *wishlistController) FetchAllWishlist(ctx echo.Context) error {

	userID, _, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestPagination{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.wishlistService.FetchAllWishlistByUserID(request, *userID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchWishlist : Fetch Wishlist by ID
// @Summary Fetch Wishlist by ID
// @Description This endpoint for fetch wishlist by ID
// @Tags Wishlist
// @Accept  json
// @Produce  json
// @Param id path string true "Wishlist Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Wishlist} "desc"
// @Router /wishlist/{id} [get]
func (c *wishlistController) FetchWishlist(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.wishlistService.FetchWishlistQuery(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteWishlist : Delete Wishlist by ID
// @Summary Delete Wishlist by ID
// @Description This endpoint for Delete wishlist
// @Tags Wishlist
// @Accept  json
// @Produce  json
// @Param id path string true "Wishlist Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Wishlist} "desc"
// @Router /wishlist/{id} [DELETE]
func (c *wishlistController) DeleteWishlist(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := uuid.Parse(parameter)
	if err != nil {
		return err
	}

	responses, err := c.wishlistService.DeleteWishlistCommand(converted, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
