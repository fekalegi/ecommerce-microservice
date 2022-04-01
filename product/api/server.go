package api

import (
	"ecommerce-microservice/product/api/controller"
	"ecommerce-microservice/product/api/handler"
	external "ecommerce-microservice/product/infra"
	"ecommerce-microservice/product/repository/postgres"
	"ecommerce-microservice/product/usecase/category"
	"ecommerce-microservice/product/usecase/product"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"

	_ "ecommerce-microservice/product/docs"
)

// Server Struct
type Server struct {
	Route *echo.Echo
}

// NewServer : Create Server Instance
func NewServer(e *echo.Echo) *Server {
	return &Server{
		e,
	}
}

func (server *Server) InitializeServer() {

	newDB := external.NewGormDB()

	repoCategory := postgres.NewCategoryRepository(newDB)
	categoryUC := category.NewCategoryService(repoCategory)
	categoryController := controller.NewCategoryController(categoryUC)

	repoProduct := postgres.NewProductRepository(newDB)
	productUC := product.NewProductService(repoProduct)
	productController := controller.NewProductController(productUC)

	apiGroup := server.Route.Group("/api")

	categoryController.Route(apiGroup)
	productController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5002",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
