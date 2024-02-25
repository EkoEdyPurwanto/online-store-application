package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	ProductsRepository interface {
		Save(products model.Products) error
		FindByType(categoryType string) ([]model.Products, error)
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

func (p *productsRepository) FindByType(categoryType string) ([]model.Products, error) {
	SQL := `SELECT 
                p.id,
                p.user_id,
                p.product_name,
                p.price,
                p.stock,
                p.category_id,
                p.created_at,
                p.updated_at,
                p.deleted_at
            FROM 
                products p
            JOIN 
                categories c ON p.category_id = c.id
            WHERE
                c.type = $1`

	rows, err := p.db.Query(SQL, categoryType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Products
	for rows.Next() {
		var product model.Products
		err := rows.Scan(
			&product.Id,
			&product.UserId,
			&product.ProductName,
			&product.Price,
			&product.Stock,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
