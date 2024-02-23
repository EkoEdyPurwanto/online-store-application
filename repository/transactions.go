package repository

import "database/sql"

type (
	TransactionsRepository interface {
	}

	transactionsRepository struct {
		db *sql.DB
	}
)

func NewTransactionsRepository(db *sql.DB) TransactionsRepository {
	return &transactionsRepository{
		db: db,
	}
}
