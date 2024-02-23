package usecase

import "online-store-application/repository"

type (
	ProductsUseCase interface {
	}

	productsUseCase struct {
		repo repository.ProductsRepository
	}
)

func NewProductsUseCase(repo repository.ProductsRepository) ProductsUseCase {
	return &productsUseCase{
		repo: repo,
	}
}
