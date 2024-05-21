package repository

import (
    "auth-service/models"
    "database/sql"
    "github.com/jmoiron/sqlx"
)

type AuthRepository struct {
    db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
    return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
    query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, user.Username, user.Password, user.Email)
    return err
}

func (r *AuthRepository) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    query := `SELECT id, username, password, email FROM users WHERE username=$1`
    err := r.db.Get(&user, query, username)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}
