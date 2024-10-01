package model

import (
	"database/sql"
	"time"
)

type Order struct {
	ID          int64        `db:"id"`
	UserID      int64        `db:"user_id"`
	StoreID     int64        `db:"store_id"`
	TotalPrice  int64        `db:"total_price"`
	TotalAmount int64        `db:"total_amount"`
	Status      string       `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type CreateOrderRequest struct {
	UserID   int64                  `json:"-"`
	StoreID  int64                  `json:"store_id" validate:"required"`
	Products []ProductOrdersRequest `json:"products" validate:"required"`
}

func (CreateOrderRequest) ErrorMessages(name string) map[string]string {
	return MapErrorRequest{
		"store_id": {
			"required": "StoreID field is required",
		},
		"products": {
			"reqeuired": "Products field is required",
		},
	}[name]
}

func (CreateOrderRequest) FieldName(name string) string {
	return map[string]string{
		"StoreID":  "store_id",
		"Products": "products",
	}[name]
}

type ProductOrdersRequest struct {
	ProductID int64 `json:"product_id"`
	Total     int64 `json:"total"`
	Price     int64 `json:"price"`
}

type OrderItemRequest struct {
	ProductID int64 `json:"product_id"`
	Amount    int64 `json:"amount"`
	Price     int64 `json:"price"`
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
