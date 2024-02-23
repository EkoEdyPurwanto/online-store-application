package manager

import "online-store-application/usecase"

type UseCaseManager interface {
	UsersAccountUC() usecase.UsersUseCase
	WalletsUC() usecase.WalletsUseCase
	ProductsUC() usecase.ProductsUseCase
	Categories() usecase.WalletsUseCase
	Carts() usecase.CartsUseCase
	Transactions() usecase.TransactionsUseCase
}

type useCaseManager struct {
	repositoryManager RepositoryManager
}

func (u *useCaseManager) UsersAccountUC() usecase.UsersUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) WalletsUC() usecase.WalletsUseCase {
	//TODO implement me
	panic("implement me")
}

// Constructor
func NewUseCaseManager(repositoryManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repositoryManager: repositoryManager,
	}
}

// implement interface here
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