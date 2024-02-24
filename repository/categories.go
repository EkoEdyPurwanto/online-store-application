package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	CategoriesRepository interface {
		Save(categories model.Categories) error
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

func (p *categoriesRepository) Save(categories model.Categories) error {
	SQL := `INSERT INTO categories(id, type, created_at, updated_at, deleted_at) VALUES($1, $2, $3, $4, $5)`
	_, err := p.db.Exec(SQL,
		categories.Id,
		categories.Type,
		categories.CreatedAt,
		categories.UpdatedAt,
		categories.DeletedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
