package api

import (
	"ecommerce-microservice/product/api/controller"
	"ecommerce-microservice/product/api/handler"
	"ecommerce-microservice/product/common/models"
	external "ecommerce-microservice/product/infra"
	backgrounds "ecommerce-microservice/product/infra/background"
	"ecommerce-microservice/product/repository/postgres"
	"ecommerce-microservice/product/usecase/category"
	"ecommerce-microservice/product/usecase/product"
	"ecommerce-microservice/product/usecase/wishlist"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"strings"
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

	repoWishlist := postgres.NewWishlistRepository(newDB)
	wishlistUC := wishlist.NewWishlistRepository(repoWishlist)
	wishlistController := controller.NewWishlistController(wishlistUC)

	apiGroup := server.Route.Group("/api")

	categoryController.Route(apiGroup)
	productController.Route(apiGroup)
	wishlistController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5002",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	kafkaHost := os.Getenv("ORDER_BROKER")
	arr := strings.Split(kafkaHost, ",")
	chanOrderHandler := make(chan models.KafkaNewOrderEventDto, 1000)
	chanNotif := make(chan models.KafkaNewOrderEventDto, 1000)
	go backgrounds.StartConsumerOrder(arr, os.Getenv("KAFKA_ORDER_TOPIC"), chanOrderHandler, chanNotif)
	go backgrounds.StartConsumerOrderHandler(chanNotif, productUC)

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
