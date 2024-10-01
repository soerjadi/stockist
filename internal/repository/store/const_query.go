package store

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	insertStore *sqlx.Stmt
	getByID     *sqlx.Stmt
}

const (
	insertStore = `
	INSERT INTO stores (
		name,
		description,
		address,
		manager_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING
	 	id,
		name,
		description,
		address,
		manager_id,
		created_at
	`

	getByID = `
	SELECT
		id,
		name,
		description,
		address,
		manager_id,
		created_at
	FROM
		stores
	WHERE
		id = $1
	`
)
