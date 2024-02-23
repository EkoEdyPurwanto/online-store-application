package repository

import "database/sql"

type (
	CartsRepository interface {
	}

	cartsRepository struct {
		db *sql.DB
	}
)

func NewCartsRepository(db *sql.DB) CartsRepository {
	return &cartsRepository{
		db: db,
	}
}
