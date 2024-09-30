package model

import (
	"database/sql"
	"time"
)

type Payment struct {
	ID            int64        `db:"id"`
	OrderID       int64        `db:"order_id"`
	InvoiceNumber string       `db:"invoice_number"`
	PaymentMethod string       `db:"payment_method"`
	Status        string       `db:"status"`
	ExpiredAt     time.Time    `db:"expired_at"`
	CreatedAt     time.Time    `db:"created_at"`
	PaidAt        sql.NullTime `db:"paid_at"`
}
