package repository

import (
	"database/sql"
)

type IUserRepository interface {
	GetUserByEmail(storedPassword *string, userId *int, email string) error
	SearchUserByEmail(count *int, email string) error
	CreateUser(email string, hash string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(storedPassword *string, userId *int, email string) error {
	if err := ur.db.QueryRow("SELECT id, password FROM user WHERE email = ?", email).Scan(userId, storedPassword); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) SearchUserByEmail(count *int, email string) error {
	if err := ur.db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", email).Scan(count); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(email string, hash string) error {
	if _, err := ur.db.Exec("INSERT INTO user (email, password) VALUES (?, ?)", email, hash); err != nil {
		return err
	}
	return nil
}
