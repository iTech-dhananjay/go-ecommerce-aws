package repository

import (
    "cart-service/models"
    "database/sql"
    "github.com/jmoiron/sqlx"
)

type CartRepository struct {
    db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
    return &CartRepository{db: db}
}

func (r *CartRepository) AddToCart(item *models.CartItem) error {
    query := `INSERT INTO cart (user_id, product_id, quantity) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, item.UserID, item.ProductID, item.Quantity)
    return err
}

func (r *CartRepository) GetCart(userID int) ([]*models.CartItem, error) {
    var items []*models.CartItem
    query := `SELECT id, user_id, product_id, quantity FROM cart WHERE user_id=$1`
    err := r.db.Select(&items, query, userID)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return items, err
}

func (r *CartRepository) RemoveFromCart(userID, itemID int) error {
    query := `DELETE FROM cart WHERE user_id=$1 AND id=$2`
    _, err := r.db.Exec(query, userID, itemID)
    return err
}
