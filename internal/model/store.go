package model

import "time"

type Stores struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Address     string    `db:"address"`
	ManagerID   int64     `db:"manager_id"` // related to user
	CreatedAt   time.Time `db:"created_at"`
}
