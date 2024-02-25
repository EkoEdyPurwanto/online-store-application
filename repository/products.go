package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	ProductsRepository interface {
		Save(products model.Products) error
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

func (p *productsRepository) Save(products model.Products) error {
	SQL := `INSERT INTO products(id, user_id, product_name, price, stock, category_id, created_at, updated_at, deleted_at ) VALUES ($1, $2, $3, $4,$5, $6, $7, $8, $9)`
	_, err := p.db.Exec(SQL,
		&products.Id,
		&products.UserId,
		&products.ProductName,
		&products.Price,
		&products.Stock,
		&products.CategoryId,
		&products.CreatedAt,
		&products.UpdatedAt,
		&products.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
