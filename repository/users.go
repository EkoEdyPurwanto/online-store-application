package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	UsersRepository interface {
		Save(users model.Users) error
	}

	usersRepository struct {
		db *sql.DB
	}
)

func (u *usersRepository) Save(users model.Users) error {
	SQL := `INSERT INTO users(id, username, password, email, phone_number, user_status, role, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := u.db.Exec(SQL,
		users.Id,
		users.Username,
		users.Password,
		users.Email,
		users.PhoneNumber,
		users.UserStatus,
		users.Role,
		users.CreatedAt,
		users.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}
