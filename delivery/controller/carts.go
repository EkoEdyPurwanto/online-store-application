package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-store-application/delivery/middleware"
	"online-store-application/helper"
	"online-store-application/model/dto/req"
	"online-store-application/model/dto/resp"
	"online-store-application/usecase"
)

type CartsController struct {
	cartsUC usecase.CartsUseCase
	engine     *echo.Echo
}

// Constructor
func NewCartsController(cartsUC usecase.CartsUseCase, engine *echo.Echo) *CartsController {
	return &CartsController{
		cartsUC: cartsUC,
		engine:     engine,
	}
}

func (c *CartsController) CartsRoute() {
	rg := c.engine.Group("/api/v1", middleware.AuthMiddleware)

	rg.POST("/carts", c.createCartsHandler)
}

func (c *CartsController) createCartsHandler(ctx echo.Context) error {
	var payload req.CreateCartsRequest
	err := helper.ReadFromRequestBody(ctx, &payload)
	if err != nil {
		return err
	}

	err = c.cartsUC.CreateCarts(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully save product into cart",
		Data:    nil,
	}

	return helper.WriteToResponseBody(ctx, response)
}