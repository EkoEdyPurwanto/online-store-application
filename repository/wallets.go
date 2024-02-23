package repository

import "database/sql"

type (
	WalletsRepository interface {
	}

	walletsRepository struct {
		db *sql.DB
	}
)

func NewWalletsRepository(db *sql.DB) WalletsRepository {
	return &walletsRepository{
		db: db,
	}
}
