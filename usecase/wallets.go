package usecase

import "online-store-application/repository"

type (
	WalletsUseCase interface {
	}

	walletsUseCase struct {
		repo repository.WalletsRepository
	}
)

func NewWalletsUseCase(repo repository.WalletsRepository) WalletsUseCase {
	return &walletsUseCase{
		repo: repo,
	}
}
