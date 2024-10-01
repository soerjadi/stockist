package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int64        `db:"id" json:"id"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	Weight      int64        `db:"weight" json:"weight"`
	Price       int64        `db:"price" json:"price"`
	StoreID     int64        `db:"store_id" json:"store_id"`
	Stock       int64        `db:"stock" json:"stock"`
	Images      string       `db:"images" json:"images"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at" json:"-"`
}

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Weight      int64  `json:"weight" validate:"required,numeric"`
	Price       int64  `json:"price" validate:"required,numeric"`
	StoreID     int64  `json:"store_id" validate:"required"`
	Stock       int64  `json:"stock" validate:"required,numeric"`
	Images      string `json:"images" validate:"required"`
}

func (CreateProductRequest) ErrorMessages(name string) map[string]string {
	return MapErrorRequest{
		"name": {
			"required": "Name field is required",
		},
		"description": {
			"required": "Description field is required",
		},
		"weight": {
			"required": "Weight field is required",
			"numeric":  "Wrong weight format, only numeric",
		},
		"price": {
			"required": "Price field is required",
			"numeric":  "Wrong price format, only numeric",
		},
		"store_id": {
			"required": "Store ID field is required",
		},
		"stock": {
			"required": "Stock field is required",
			"numeric":  "Wrong stock format, only numeric",
		},
		"images": {
			"required": "Images field is required",
		},
	}[name]
}

func (CreateProductRequest) FieldName(name string) string {
	return map[string]string{
		"Name":        "name",
		"Description": "description",
		"Weight":      "weight",
		"Price":       "price",
		"StoreID":     "store_id",
		"Stock":       "stock",
		"Images":      "images",
	}[name]
}
