package user

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	insertUser           *sqlx.Stmt
	getUserByID          *sqlx.Stmt
	getUserByPhoneNumber *sqlx.Stmt
	getUserByEmail       *sqlx.Stmt
}

const (
	insertUser = `
	INSERT INTO users (
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt
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
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	`

	getUserByID = `
	SELECT
		id,
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	FROM
		users
	WHERE
		id = $1
	`

	getUserByPhoneNumber = `
	SELECT
		id,
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	FROM
		users
	WHERE
		phone_number = $1
	`

	getUserByEmail = `
	SELECT
		id,
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	FROM
		users
	WHERE
		email = $1
	`
)
