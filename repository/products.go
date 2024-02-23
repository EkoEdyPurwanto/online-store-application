package repository

import "database/sql"

type (
	ProductsRepository interface {
	}

	productsRepository struct {
		db *sql.DB
	}
)

func NewProductsRepository(db *sql.DB) ProductsRepository {
	return &productsRepository{
		db: db,
	}
}
