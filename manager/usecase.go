package manager

import "online-store-application/usecase"

type UseCaseManager interface {
	UsersUC() usecase.UsersUseCase
	WalletsUC() usecase.WalletsUseCase
	ProductsUC() usecase.ProductsUseCase
	CategoriesUC() usecase.CategoriesUseCase
	CartsUC() usecase.CartsUseCase
	TransactionsUC() usecase.TransactionsUseCase
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
	return usecase.NewProductsUseCase(u.repositoryManager.ProductsRepository())
}

func (u *useCaseManager) CategoriesUC() usecase.CategoriesUseCase {
	return usecase.NewCategoriesUseCase(u.repositoryManager.CategoriesRepository())
}

func (u *useCaseManager) CartsUC() usecase.CartsUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) TransactionsUC() usecase.TransactionsUseCase {
	//TODO implement me
	panic("implement me")
}
