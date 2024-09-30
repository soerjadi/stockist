package model

import (
	"database/sql"
	"time"
)

type Order struct {
	ID          int64        `db:"id"`
	UserID      int64        `db:"user_id"`
	ShopID      int64        `db:"shop_id"`
	TotalPrice  int64        `db:"total_price"`
	TotalAmount int64        `db:"total_amount"`
	Status      string       `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type OrderItem struct {
	ID        int64     `db:"id"`
	OrderID   int64     `db:"order_id"`
	ProductID int64     `db:"product_id"`
	Amount    int64     `db:"amount"`
	Price     int64     `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

const (
	ORDER_STATUS_PAID     = "paid"
	ORDER_STATUS_UNPAID   = "unpaid"
	ORDER_STATUS_CREATED  = "created"
	ORDER_STATUS_STALE    = "stale" // stale is when customer already checkout but not yet confirmed the payment
	ORDER_STATUS_DELIVERY = "delivery"
)
