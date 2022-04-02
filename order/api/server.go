package api

import (
	"ecommerce-microservice/order/api/handler"
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
	//
	//newDB := external.NewGormDB()

	//repoWishlist := postgres.NewWishlistRepository(newDB)
	//wishlistUC := wishlist.NewWishlistRepository(repoWishlist)
	//wishlistController := controller.NewWishlistController(wishlistUC)
	//
	//apiGroup := server.Route.Group("/api")
	//
	//wishlistController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5002",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
