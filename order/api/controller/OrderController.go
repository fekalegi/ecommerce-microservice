package controller

import (
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/common/helper"
	"ecommerce-microservice/order/common/models"
	"ecommerce-microservice/order/usecase"
	"github.com/fekalegi/custom-package/authentications/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

// OrderController : Controller for Order Api
type orderController struct {
	orderService usecase.OrderService
}

// NewOrderController : Instance for register  OrderService
func NewOrderController(orderService usecase.OrderService) *orderController {
	return &orderController{orderService: orderService}
}

// Route : Setting Order Route
func (c *orderController) Route(group *echo.Group) {
	group.POST("/order/:id/products", c.AddOrUpdateOrderProducts, middleware.KeyAuth(middlewares.AuthCheck))
	group.POST("/order/:id/rating", c.AddRatingSeller, middleware.KeyAuth(middlewares.AuthCheck))
	group.POST("/order", c.CreateInitialOrder, middleware.KeyAuth(middlewares.AuthCheck))
	group.DELETE("/order/:id", c.DeleteOrder, middleware.KeyAuth(middlewares.AuthCheck))
	group.DELETE("/order/product/:id", c.DeleteOrderProduct, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/order", c.FetchAllOrder, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/order/:id", c.FetchOrder, middleware.KeyAuth(middlewares.AuthCheck))
	group.GET("/order/seller/:id", c.FetchRatingSeller, middleware.KeyAuth(middlewares.AuthCheck))
	group.PATCH("/order/:id/pay", c.UpdateStatusOrderToPayment, middleware.KeyAuth(middlewares.AuthCheck))
	group.PATCH("/order/:id/pack", c.UpdateStatusOrderToPacking, middleware.KeyAuth(middlewares.AuthCheck))
	group.PATCH("/order/:id/shipped", c.UpdateStatusOrderToShipped, middleware.KeyAuth(middlewares.AuthCheck))
	group.PATCH("/order/:id/arrived", c.UpdateStatusOrderToArrived, middleware.KeyAuth(middlewares.AuthCheck))
	group.PATCH("/order/:id/finished", c.UpdateStatusOrderToFinished, middleware.KeyAuth(middlewares.AuthCheck))
}

// UpdateStatusOrderToFinished godoc
// @Summary UpdateStatusOrderToFinished
// @Description This endpoint for UpdateStatusOrderToFinished
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/finished [patch]
func (c *orderController) UpdateStatusOrderToFinished(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.UpdateStatusOrderToFinished(int64(converted), *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateStatusOrderToArrived godoc
// @Summary UpdateStatusOrderToArrived
// @Description This endpoint for UpdateStatusOrderToArrived
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/arrived [patch]
func (c *orderController) UpdateStatusOrderToArrived(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.UpdateStatusOrderToItemsHaveArrived(int64(converted), *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateStatusOrderToShipped godoc
// @Summary UpdateStatusOrderToShipped
// @Description This endpoint for UpdateStatusOrderToShipped
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/shipped [patch]
func (c *orderController) UpdateStatusOrderToShipped(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.UpdateStatusOrderToItemsBeingShipped(int64(converted), *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateStatusOrderToPacking godoc
// @Summary UpdateStatusOrderToPacking
// @Description This endpoint for UpdateStatusOrderToPacking
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/pack [patch]
func (c *orderController) UpdateStatusOrderToPacking(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.UpdateStatusOrderToPackingItems(int64(converted), *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// UpdateStatusOrderToPayment godoc
// @Summary UpdateStatusOrderToPayment
// @Description This endpoint for UpdateStatusOrderToPayment
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/pay [patch]
func (c *orderController) UpdateStatusOrderToPayment(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.UpdateStatusOrderToWaitForPayment(int64(converted), *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// CreateInitialOrder godoc
// @Summary CreateInitialOrder New Order
// @Description This endpoint for CreateInitialOrder
// @Tags Order
// @Accept  json
// @Produce  json
// @Param services body dto.RequestInitialOrder true "Create new order"
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order [post]
func (c *orderController) CreateInitialOrder(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestInitialOrder{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.orderService.CreateInitialOrderCommand(request, *userID, *level)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// AddOrUpdateOrderProducts godoc
// @Summary AddOrUpdateOrderProducts New Order
// @Description This endpoint for AddOrUpdateOrderProducts
// @Tags Order
// @Accept  json
// @Produce  json
// @Param services body dto.RequestOrderProduct true "Add or Update Order Product"
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/products [post]
func (c *orderController) AddOrUpdateOrderProducts(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	userID, _, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestOrderProduct{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.orderService.AddOrUpdateOrderProductsCommand(request, int64(converted), *userID)

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// AddRatingSeller godoc
// @Summary AddRatingSeller
// @Description This endpoint for AddRatingSeller
// @Tags Order
// @Accept  json
// @Produce  json
// @Param services body dto.RequestRatingSeller true "Request Rating"
// @Param id path string true "ID Order"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /order/{id}/rating [post]
func (c *orderController) AddRatingSeller(ctx echo.Context) error {
	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	userID, _, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestRatingSeller{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.orderService.AddRatingSeller(request, *userID, int64(converted))

	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchAllOrder : Fetch All Order
// @Summary Fetch All Order
// @Description This endpoint for fetch all order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param services query dto.RequestFetchOrderByStatus true "Parameter with pagination"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=[]domain.Order} "desc"
// @Router /order [get]
func (c *orderController) FetchAllOrder(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	request := dto.RequestFetchOrderByStatus{}

	err = models.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.orderService.FetchAllOrderQuery(request, *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchOrder : Fetch Order by ID
// @Summary Fetch Order by ID
// @Description This endpoint for fetch order by ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Order} "desc"
// @Router /order/{id} [get]
func (c *orderController) FetchOrder(ctx echo.Context) error {

	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.FetchOrderQuery(int64(converted), *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// FetchRatingSeller : FetchRatingSeller by ID
// @Summary FetchRatingSeller by ID
// @Description This endpoint for fetch order by ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Order} "desc"
// @Router /order/seller/{id} [get]
func (c *orderController) FetchRatingSeller(ctx echo.Context) error {

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.FetchRatingSellerBySellerID(converted)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteOrder : Delete Order by ID
// @Summary Delete Order by ID
// @Description This endpoint for Delete order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order Id"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Order} "desc"
// @Router /order/{id} [DELETE]
func (c *orderController) DeleteOrder(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.DeleteOrderCommand(int64(converted), *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// DeleteOrderProduct : Delete Order Product by ID
// @Summary Delete Order Product by ID
// @Description This endpoint for Delete order Product
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order Product ID"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed{data=domain.Order} "desc"
// @Router /order/product/{id} [DELETE]
func (c *orderController) DeleteOrderProduct(ctx echo.Context) error {
	userID, level, err := helper.GetUserDetailFromContext(ctx)

	if err != nil {
		return err
	}

	parameter := ctx.Param("id")

	converted, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	responses, err := c.orderService.DeleteOrderProductCommand(int64(converted), *userID, *level)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
