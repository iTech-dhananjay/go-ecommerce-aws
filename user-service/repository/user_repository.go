package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"user-service/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Email)
	return err
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, email FROM users WHERE id=$1`
	err := r.db.Get(&user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username=$1, password=$2, email=$3 WHERE id=$4`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Email, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}
