package main

import (
	"ecommerce-microservice/user/api"
	"fmt"
	"github.com/labstack/echo/v4"
)

// @title Mini Rest Api
// @version 1.0
// @description This is swagger for  Mini Rest Api
// @termsOfService http://swagger.io/terms/

// @contact.name Feka Legi Heryana Rizki
// @contact.email fekalegi@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	fmt.Println("application has started...")
	newEcho := echo.New()
	serverAPI := api.NewServer(newEcho)

	serverAPI.InitializeServer()
}
