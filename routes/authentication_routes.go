package routes

import (
	"genesis_auth/controller"

	"github.com/gin-gonic/gin"
)

type AuthenticationRoutes interface {
	Routes(*gin.Engine)
}

type route struct {
	controller controller.AuthenticationController
}

func NewAuthecationRoutes(controller controller.AuthenticationController) AuthenticationRoutes {
	return &route{
		controller: controller,
	}
}

func (r *route) Routes(router *gin.Engine) {
	//all routes related to contributions
	router.POST("/user/create", r.controller.CreateUser())
	router.POST("/user/login", r.controller.LogIn())
}
