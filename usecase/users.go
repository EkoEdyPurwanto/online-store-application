package usecase

import "online-store-application/repository"

type (
	UsersUseCase interface {
	}

	usersUseCase struct {
		repo repository.UsersRepository
	}
)

func NewUsersUseCase(repo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{
		repo: repo,
	}
}
