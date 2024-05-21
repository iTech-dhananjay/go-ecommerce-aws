package models

type Payment struct {
    ID        int    `db:"id"`
    OrderID   int    `db:"order_id"`
    Amount    int    `db:"amount"`
    Status    string `db:"status"`
}
