package model

import (
	"database/sql"
	"time"
)

type Warehouse struct {
	ID              int64        `db:"id"`
	Name            string       `db:"name"`
	Address         string       `db:"address"`
	Status          string       `db:"status"`
	Quota           int64        `db:"quota"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
	StatusUpdatedAt sql.NullTime `db:"status_updated_at"`
}

const (
	WAREHOUSE_STATUS_ACTIVE   = "active"
	WAREHOUSE_STATUS_INACTIVE = "invactive"
)

func (w Warehouse) IsActive() bool {
	return w.Status == WAREHOUSE_STATUS_ACTIVE
}
