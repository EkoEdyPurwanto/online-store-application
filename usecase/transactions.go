package usecase

import "online-store-application/repository"

type (
	TransactionsUseCase interface {
	}

	transactionsUseCase struct {
		repo repository.TransactionsRepository
	}
)

func NewTransactionsUseCase(repo repository.TransactionsRepository) TransactionsUseCase {
	return &transactionsUseCase{
		repo: repo,
	}
}
