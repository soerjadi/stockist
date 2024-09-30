package model

import (
	"database/sql"
	"time"
)

type WarehouseProduct struct {
	ID          int64        `db:"id"`
	WarehouseID int64        `db:"warehouse_id"`
	ProductID   int64        `db:"product_id"`
	Stock       int64        `db:"stock"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type WarehouseTrfLog struct {
	ID          int64        `db:"id"`
	TargetID    int64        `db:"target_id"`
	TargetClass string       `db:"target_class"`
	ProductID   int64        `db:"product_id"`
	Amount      int64        `db:"amount"`
	Note        string       `db:"note"`
	Status      string       `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

const (
	WAREHOUSE_TARGET_CLASS_WAREHOUSE = "__warehouse__" // it means transfer to another warehouse
	WAREHOUSE_TARGET_CLASS_USER      = "__user__"      // it means user order this product

	WAREHOUSE_TRANSFER_STATUS_RECEIVED         = "received"
	WAREHOUSE_TRANSFER_STATUS_INPROGRESS       = "inprogress"
	WAREHOUSE_TRANSFER_STATUS_PROCESS_TRANSFER = "process_transfer"
	WAREHOUSE_TRANSFER_STATUS_TRANSFERED       = "transfered"
)
