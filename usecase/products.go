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
	ProductsUseCase interface {
		CreateProducts(payload req.CreateProductsRequest) error
		GetProductsByType(payload string) ([]model.Products, error)
	}

	productsUseCase struct {
		repo repository.ProductsRepository
	}
)

// Constructor
func NewProductsUseCase(repo repository.ProductsRepository) ProductsUseCase {
	return &productsUseCase{
		repo: repo,
	}
}

func (p *productsUseCase) CreateProducts(payload req.CreateProductsRequest) error {
	// Validate the payload
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}

	products := model.Products{
		Id:          common.GenerateID(),
		UserId:      payload.UserId,
		ProductName: payload.ProductName,
		Price:       payload.Price,
		Stock:       payload.Stock,
		CategoryId:  payload.CategoryId,
		CreatedAt:   time.Now(),
	}

	err = p.repo.Save(products)
	if err != nil {
		return fmt.Errorf("failed save product: %v", err.Error())
	}

	return nil
}

func (p *productsUseCase) GetProductsByType(payload string) ([]model.Products, error) {
	getProductBytype, err := p.repo.FindByType(payload)
	if err != nil {
		return []model.Products{}, err
	}

	return getProductBytype, nil
}