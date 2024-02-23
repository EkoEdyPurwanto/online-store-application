package manager

import "online-store-application/usecase"

type UseCaseManager interface {
	UsersUC() usecase.UsersUseCase
	WalletsUC() usecase.WalletsUseCase
	ProductsUC() usecase.ProductsUseCase
	Categories() usecase.WalletsUseCase
	Carts() usecase.CartsUseCase
	Transactions() usecase.TransactionsUseCase
}

type useCaseManager struct {
	repositoryManager RepositoryManager
}

// Constructor
func NewUseCaseManager(repositoryManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repositoryManager: repositoryManager,
	}
}

// implement interface here
func (u *useCaseManager) UsersUC() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repositoryManager.UsersRepository(), u.WalletsUC())
}

func (u *useCaseManager) WalletsUC() usecase.WalletsUseCase {
	return usecase.NewWalletsUseCase(u.repositoryManager.WalletsRepository())
}

func (u *useCaseManager) ProductsUC() usecase.ProductsUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) Categories() usecase.WalletsUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) Carts() usecase.CartsUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) Transactions() usecase.TransactionsUseCase {
	//TODO implement me
	panic("implement me")
}
