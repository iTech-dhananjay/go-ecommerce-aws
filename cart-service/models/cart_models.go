package models

type CartItem struct {
    ID       int `db:"id"`
    UserID   int `db:"user_id"`
    ProductID int `db:"product_id"`
    Quantity int `db:"quantity"`
}
