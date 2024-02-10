package main

import (
	"fmt"
	"genesis_auth/config"
	"genesis_auth/controller"
	"genesis_auth/routes"
	"genesis_auth/service"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	authenticationService    service.AuthenticationService       = service.NewAuthenticationService(config.GetCollection(config.DB, os.Getenv("AUTH_STAGING")))
	authenticationController controller.AuthenticationController = controller.NewContributionController(authenticationService)
	authenticationRoutes     routes.AuthenticationRoutes         = routes.NewAuthecationRoutes(authenticationController)
)

func main() {
	router := gin.New()

	//database
	config.ConnectToDB()

	router.Use(gin.Logger())

	authenticationRoutes.Routes(router)

	fmt.Println("Environment variable: ", os.Getenv("AUTH_STAGING"))


	router.Run()
}
