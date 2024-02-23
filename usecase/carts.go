package usecase

import "online-store-application/repository"

type (
	CartsUseCase interface {
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
