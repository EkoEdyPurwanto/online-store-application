package repository

import "database/sql"

type (
	 CategoriesRepository interface {
	}

	categoriesRepository struct {
		db *sql.DB
	}
)

func NewCategoriesRepository(db *sql.DB) CategoriesRepository {
	return &categoriesRepository{
		db: db,
	}
}

