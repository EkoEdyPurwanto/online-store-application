package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-store-application/helper"
	"online-store-application/model/dto/req"
	"online-store-application/model/dto/resp"
	"online-store-application/usecase"
)

type UsersController struct {
	usersUC   usecase.UsersUseCase
	walletsUC usecase.WalletsUseCase
	engine    *echo.Echo
}

// Constructor
func NewUsersController(usersUC usecase.UsersUseCase, walletUC usecase.WalletsUseCase, engine *echo.Echo) *UsersController {
	return &UsersController{
		usersUC:   usersUC,
		walletsUC: walletUC,
		engine:    engine,
	}
}

func (ua *UsersController) AuthRoute() {
	rg := ua.engine.Group("/api/v1")

	rg.POST("/auth/register", ua.registerHandler)
}

func (ua *UsersController) registerHandler(c echo.Context) error {
	var payload req.RegisterRequest
	err := helper.ReadFromRequestBody(c, &payload)
	if err != nil {
		return err
	}

	err = ua.usersUC.Register(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully register",
		Data:    nil,
	}

	return helper.WriteToResponseBody(c, response)
}
