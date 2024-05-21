package repository

import (
    "payment-service/models"
    "database/sql"
    "github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
    db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
    return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment *models.Payment) error {
    query := `INSERT INTO payments (order_id, amount, status) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, payment.OrderID, payment.Amount, payment.Status)
    return err
}

func (r *PaymentRepository) GetPaymentByID(id int) (*models.Payment, error) {
    var payment models.Payment
    query := `SELECT id, order_id, amount, status FROM payments WHERE id=$1`
    err := r.db.Get(&payment, query, id)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &payment, err
}

func (r *PaymentRepository) UpdatePayment(payment *models.Payment) error {
    query := `UPDATE payments SET order_id=$1, amount=$2, status=$3 WHERE id=$4`
    _, err := r.db.Exec(query, payment.OrderID, payment.Amount, payment.Status, payment.ID)
    return err
}

func (r *PaymentRepository) DeletePayment(id int) error {
    query := `DELETE FROM payments WHERE id=$1`
    _, err := r.db.Exec(query, id)
    return err
}
