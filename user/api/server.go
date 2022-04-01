package api

import (
	"ecommerce-microservice/user/api/controller"
	"ecommerce-microservice/user/api/handler"
	"ecommerce-microservice/user/common/helper"
	external "ecommerce-microservice/user/infra"
	"ecommerce-microservice/user/repository/postgres"
	"ecommerce-microservice/user/usecase/role"
	"ecommerce-microservice/user/usecase/user"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"

	_ "ecommerce-microservice/user/docs"
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

	repoRole := postgres.NewRoleRepository(newDB)
	roleUC := role.NewRoleImplementation(repoRole)
	roleController := controller.NewRoleController(roleUC)

	apiGroup := server.Route.Group("/api")

	userController.Route(apiGroup)
	roleController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5000",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
