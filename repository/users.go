package repository

import "database/sql"

type (
	UsersRepository interface {
	}

	usersRepository struct {
		db *sql.DB
	}
)

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}
