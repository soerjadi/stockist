package product

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getByID       *sqlx.Stmt
	getList       *sqlx.Stmt
	createProduct *sqlx.Stmt
}

const (
	getByID = `
	SELECT
		id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	FROM 
		products
	WHERE 
		id = $1
	`

	getList = `
	SELECT
		id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	FROM
		products
	ORDER BY id DESC
	LIMIT $1
	OFFSET $2
	`

	createProduct = `
	INSERT INTO products (
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	) RETURNING
	 	id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images,
		created_at
	`
)
