package repository

import (
    "order-service/models"
    "database/sql"
    "github.com/jmoiron/sqlx"
)

type OrderRepository struct {
    db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
    return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
    query := `INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(query, order.UserID, order.ProductID, order.Quantity, order.Status)
    return err
}

func (r *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
    var order models.Order
    query := `SELECT id, user_id, product_id, quantity, status FROM orders WHERE id=$1`
    err := r.db.Get(&order, query, id)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &order, err
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
    query := `UPDATE orders SET user_id=$1, product_id=$2, quantity=$3, status=$4 WHERE id=$5`
    _, err := r.db.Exec(query, order.UserID, order.ProductID, order.Quantity, order.Status, order.ID)
    return err
}

func (r *OrderRepository) DeleteOrder(id int) error {
    query := `DELETE FROM orders WHERE id=$1`
    _, err := r.db.Exec(query, id)
    return err
}
