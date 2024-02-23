package usecase

import (
	"online-store-application/model"
	"online-store-application/repository"
)

type (
	WalletsUseCase interface {
		CreateWallet(payload model.Wallets) error
	}

	walletsUseCase struct {
		repo repository.WalletsRepository
	}
)

func (w *walletsUseCase) CreateWallet(payload model.Wallets) error {
	err := w.repo.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func NewWalletsUseCase(repo repository.WalletsRepository) WalletsUseCase {
	return &walletsUseCase{
		repo: repo,
	}
}
