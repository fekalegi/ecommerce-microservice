package api

import (
	"ecommerce-microservice/order/api/controller"
	"ecommerce-microservice/order/api/handler"
	"ecommerce-microservice/order/common/helper"
	external "ecommerce-microservice/order/infra"
	"ecommerce-microservice/order/repository/postgres"
	"ecommerce-microservice/order/usecase/note"
	"ecommerce-microservice/order/usecase/user"
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
	newDB := external.NewGormDB()

	helperInterface := helper.NewHelperFunction()
	repoUser := postgres.NewUserRepository(newDB)
	userUC := user.NewUserImplementation(repoUser, helperInterface)
	userController := controller.NewUserController(userUC)

	repoNote := postgres.NewNoteRepository(newDB)
	noteUC := note.NewNoteImplementation(repoNote)
	noteController := controller.NewNoteController(noteUC)

	apiGroup := server.Route.Group("/api")

	userController.Route(apiGroup)
	noteController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5002",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
