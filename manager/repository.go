package manager

import "online-store-application/repository"

type RepositoryManager interface {
	UsersRepository() repository.UsersRepository
	WalletsRepository() repository.WalletsRepository
	ProductsRepository() repository.ProductsRepository
	CategoriesRepository() repository.CategoriesRepository
	CartsRepository() repository.CartsRepository
	TransactionsRepository() repository.TransactionsRepository

}

type repositoryManager struct {
	infraManager InfraManager
}

// Constructor
func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}

// implement interface here
func (r *repositoryManager) UsersRepository() repository.UsersRepository {
	return repository.NewUsersRepository(r.infraManager.Conn())
}

func (r *repositoryManager) WalletsRepository() repository.WalletsRepository {
	return repository.NewWalletsRepository(r.infraManager.Conn())
}

func (r *repositoryManager) ProductsRepository() repository.ProductsRepository {
	return repository.NewProductsRepository(r.infraManager.Conn())
}

func (r *repositoryManager) CategoriesRepository() repository.CategoriesRepository {
	return repository.NewCategoriesRepository(r.infraManager.Conn())
}

func (r *repositoryManager) CartsRepository() repository.CartsRepository {
	return repository.NewCartsRepository(r.infraManager.Conn())
}

func (r *repositoryManager) TransactionsRepository() repository.TransactionsRepository {
	return repository.NewTransactionsRepository(r.infraManager.Conn())
}