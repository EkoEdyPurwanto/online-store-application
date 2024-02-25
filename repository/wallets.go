package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	WalletsRepository interface {
		Save(wallets model.Wallets) error
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

func (w *walletsRepository) Save(wallets model.Wallets) error {
	SQL := `INSERT INTO wallets(id, user_id, rekening_user, balance, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) `
	_, err := w.db.Exec(SQL,
		&wallets.Id,
		&wallets.UserId,
		&wallets.RekeningUser,
		&wallets.Balance,
		&wallets.CreatedAt,
		&wallets.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
