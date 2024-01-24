package controller

import (
	"genesis_auth/dto"
	"genesis_auth/responses"
	"genesis_auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthenticationController interface {
	CreateUser() gin.HandlerFunc
}

type authenticationController struct {
	service service.AuthenticationService
}

func NewContributionController(svc service.AuthenticationService) AuthenticationController {
	return &authenticationController{
		service: svc,
	}
}

func (controller authenticationController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDto dto.UserDto

		if err := c.BindJSON(&userDto); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthenticationResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			})
			return
		}

		//check required fields
		if validatorErr := validate.Struct(&userDto); validatorErr != nil {
			c.JSON(http.StatusBadRequest, responses.AuthenticationResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validatorErr.Error(),
			})
			return
		}

		result, addErr := controller.service.CreateUser(&userDto)

		if addErr != nil {
			c.JSON(http.StatusBadRequest, responses.AuthenticationResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    addErr.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, responses.AuthenticationResponse{
			Status:  http.StatusCreated,
			Message: "User Created",
			Data:    result,
		})
	}
}
