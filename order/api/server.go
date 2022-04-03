package api

import (
	"ecommerce-microservice/order/api/controller"
	"ecommerce-microservice/order/api/handler"
	"ecommerce-microservice/order/common/command"
	external "ecommerce-microservice/order/infra"
	"ecommerce-microservice/order/infra/messaging"
	"ecommerce-microservice/order/repository/postgres"
	"ecommerce-microservice/order/usecase/order"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"

	_ "ecommerce-microservice/order/docs"
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

	kafkaCredentials := command.GetKafkaCredentials()
	producerService, _ := messaging.NewProducer(kafkaCredentials)

	newDB := external.NewGormDB()

	repoOrder := postgres.NewOrderRepository(newDB)
	orderUC := order.NewOrderService(repoOrder, producerService)
	orderController := controller.NewOrderController(orderUC)

	apiGroup := server.Route.Group("/api")

	orderController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5001",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
