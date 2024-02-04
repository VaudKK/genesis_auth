package main

import (
	"genesis_auth/config"
	"genesis_auth/controller"
	"genesis_auth/routes"
	"genesis_auth/service"

	"github.com/gin-gonic/gin"
)

var (
	authenticationService    service.AuthenticationService       = service.NewAuthenticationService(config.GetCollection(config.DB, "gen_users"))
	authenticationController controller.AuthenticationController = controller.NewContributionController(authenticationService)
	authenticationRoutes     routes.AuthenticationRoutes         = routes.NewAuthecationRoutes(authenticationController)
)

func main() {
	router := gin.New()

	//database
	config.ConnectToDB()

	router.Use(gin.Logger())

	authenticationRoutes.Routes(router)

	router.Run()
}
