package repository

import (
	"database/sql"
	"online-store-application/model"
)

type (
	UsersRepository interface {
		Save(users model.Users) error
		FindUserByIdentifier(identifier string) (model.Users, error)
		FindById(id string) (model.Users, error)
	}

	usersRepository struct {
		db *sql.DB
	}
)

// Constructor
func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (u *usersRepository) Save(users model.Users) error {
	SQL := `INSERT INTO users(id, username, password, email, phone_number, user_status, role, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := u.db.Exec(SQL,
		&users.Id,
		&users.Username,
		&users.Password,
		&users.Email,
		&users.PhoneNumber,
		&users.UserStatus,
		&users.Role,
		&users.CreatedAt,
		&users.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepository) FindUserByIdentifier(identifier string) (model.Users, error) {
	SQL := `SELECT * FROM users WHERE (username = $1 OR email = $2 OR phone_number = $3)`
	row := u.db.QueryRow(SQL, identifier, identifier, identifier)
	var users model.Users
	err := row.Scan(
		&users.Id,
		&users.Username,
		&users.Password,
		&users.Email,
		&users.PhoneNumber,
		&users.UserStatus,
		&users.Role,
		&users.CreatedAt,
		&users.UpdatedAt,
	)

	if err != nil {
		return model.Users{}, err
	}
	return users, nil
}

func (u *usersRepository) FindById(id string) (model.Users, error) {
	SQL := `SELECT * FROM users WHERE (id = $1)`
	row := u.db.QueryRow(SQL, id)
	var users model.Users
	err := row.Scan(
		&users.Id,
		&users.Username,
		&users.Password,
		&users.Email,
		&users.PhoneNumber,
		&users.UserStatus,
		&users.Role,
		&users.CreatedAt,
		&users.UpdatedAt,
	)

	if err != nil {
		return model.Users{}, err
	}
	return users, nil
}