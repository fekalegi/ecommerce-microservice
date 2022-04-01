package api

import (
	"ecommerce-microservice/product/api/controller"
	"ecommerce-microservice/product/api/handler"
	"ecommerce-microservice/product/common/helper"
	external "ecommerce-microservice/product/infra"
	"ecommerce-microservice/product/repository/postgres"
	"ecommerce-microservice/product/usecase/note"
	"ecommerce-microservice/product/usecase/user"
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
		Addr:         ":5000",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
