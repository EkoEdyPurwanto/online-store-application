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

type ProductsController struct {
	productsUC usecase.ProductsUseCase
	engine       *echo.Echo
}

// Constructor
func NewProductsController(productsUC usecase.ProductsUseCase, engine *echo.Echo) *ProductsController {
	return &ProductsController{
		productsUC:productsUC,
		engine:       engine,
	}
}

func (p *ProductsController) ProductsRoute() {
	rg := p.engine.Group("/api/v1", middleware.AuthMiddleware)

	rg.POST("/products", p.createProductsHandler)
}

func (p *ProductsController) createProductsHandler(ctx echo.Context) error {
	var payload req.CreateProductsRequest
	err := helper.ReadFromRequestBody(ctx, &payload)
	if err != nil {
		return err
	}

	err = p.productsUC.CreateProducts(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully create product",
		Data:    nil,
	}

	return helper.WriteToResponseBody(ctx, response)
}
