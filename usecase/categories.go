package usecase

import "online-store-application/repository"

type (
	CategoriesUseCase interface {
	}

	categoriesUseCase struct {
		repo repository.CategoriesRepository
	}
)

func NewCategoriesUseCase(repo repository.CategoriesRepository) CategoriesUseCase {
	return &categoriesUseCase{
		repo: repo,
	}
}
