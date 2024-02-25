package repository

import (
	"database/sql"
	"errors"
	"online-store-application/model"
)

type (
	CartsRepository interface {
		Save(carts model.Carts) error
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

func (c *cartsRepository) Save(carts model.Carts) error {
	SQL := `INSERT INTO carts(id, user_id, product_id, quantity, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := c.db.Exec(SQL,
		&carts.Id,
		&carts.UserId,
		&carts.ProductId,
		&carts.Quantity,
		&carts.CreatedAt,
		&carts.UpdatedAt,
		&carts.DeletedAt,
	)

	if err != nil {
		return errors.New("failed to save cart: " + err.Error())
	}
	return nil
}
