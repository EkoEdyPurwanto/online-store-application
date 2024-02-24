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

func (u *UsersController) AuthRoute() {
	rg := u.engine.Group("/api/v1")

	rg.POST("/auth/register", u.registerHandler)
	rg.POST("/auth/login", u.loginHandler)
}

func (u *UsersController) registerHandler(ctx echo.Context) error {
	var payload req.RegisterRequest
	err := helper.ReadFromRequestBody(ctx, &payload)
	if err != nil {
		return err
	}

	err = u.usersUC.Register(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully register",
		Data:    nil,
	}

	return helper.WriteToResponseBody(ctx, response)
}

func (u *UsersController) loginHandler(ctx echo.Context) error {
	var payload req.LoginRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, resp.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	token, err := u.usersUC.Login(payload)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, resp.ApiResponse{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, resp.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully login",
		Data:    token,
	})
}
