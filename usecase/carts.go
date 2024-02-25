package usecase

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"online-store-application/model"
	"online-store-application/model/dto/req"
	"online-store-application/repository"
	"online-store-application/utils/common"
	"time"
)

type (
	CartsUseCase interface {
		CreateCarts(payload req.CreateCartsRequest) error
	}

	cartsUseCase struct {
		repo repository.CartsRepository
	}
)

func NewCartsUseCase(repo repository.CartsRepository) CartsUseCase {
	return &cartsUseCase{
		repo: repo,
	}
}

func (c *cartsUseCase) CreateCarts(payload req.CreateCartsRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}

	carts := model.Carts{
		Id:        common.GenerateID(),
		UserId:    payload.UserId,
		ProductId: payload.ProductId,
		Quantity:  payload.Quantity,
		CreatedAt: time.Now(),
	}

	err = c.repo.Save(carts)
	if err != nil {
		return fmt.Errorf("failed save product to cart: %v", err.Error())
	}

	return nil
}