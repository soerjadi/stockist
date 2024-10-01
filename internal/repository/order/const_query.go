package order

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	createOrder       *sqlx.Stmt
	createOrderItem   *sqlx.Stmt
	updateOrderStatus *sqlx.Stmt
}

const (
	createOrder = `
	INSERT INTO orders (
		user_id,
		store_id,
		status
	) VALUES (
		$1,
		$2,
		$3
	) RETURNING 
	 	id,
		user_id,
		store_id,
		status,
		created_at
	`

	createOrderItem = `
	INSERT INTO order_item (
		order_id,
		product_id,
		amount,
		price
	) VALUES (
		$1,
		$2,
		$3,
		$4 
	)
	`

	updateOrderStatus = `
	UPDATE 
		orders 
	SET
		status = $1
	WHERE 
		id = $2
	`
)
