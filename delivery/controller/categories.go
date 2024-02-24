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

type CategoriesController struct {
	categoriesUC usecase.CategoriesUseCase
	engine       *echo.Echo
}

// Constructor
func NewCategoriesController(categoriesUC usecase.CategoriesUseCase, engine *echo.Echo) *CategoriesController {
	return &CategoriesController{
		categoriesUC: categoriesUC,
		engine:       engine,
	}
}

func (c *CategoriesController) CategoriesRoute() {
	rg := c.engine.Group("/api/v1", middleware.AuthMiddleware)

	rg.POST("/categories", c.createCategoriesHandler)
}

func (c *CategoriesController) createCategoriesHandler(ctx echo.Context) error {
	var payload req.CreateCategoriesRequest
	err := helper.ReadFromRequestBody(ctx, &payload)
	if err != nil {
		return err
	}

	err = c.categoriesUC.CreateCategories(payload)
	if err != nil {
		return err
	}

	response := resp.ApiResponse{
		Status:  http.StatusCreated,
		Message: "successfully create category",
		Data:    nil,
	}

	return helper.WriteToResponseBody(ctx, response)
}
