package model

type Product struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Weight      int64  `db:"weight"`
	Price       int64  `db:"price"`
	StoreID     int64  `db:"store_id"`
	Stock       int64  `db:"stock"`
	Images      string `db:"images"`
}
