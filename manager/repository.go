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
	//TODO implement me
	panic("implement me")
}

func (r *repositoryManager) WalletsRepository() repository.WalletsRepository {
	return repository.NewWalletsRepository(r.infraManager.Conn())
}

func (r *repositoryManager) ProductsRepository() repository.ProductsRepository {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryManager) CategoriesRepository() repository.CategoriesRepository {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryManager) CartsRepository() repository.CartsRepository {
	//TODO implement me
	panic("implement me")
}

func (r *repositoryManager) TransactionsRepository() repository.TransactionsRepository {
	//TODO implement me
	panic("implement me")
}