package model

import "time"

type Store struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Address     string    `db:"address" json:"address"`
	ManagerID   int64     `db:"manager_id" json:"-"` // related to user
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type RegisterStoreRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	OwnerID     int64
}

func (RegisterStoreRequest) ErrorMessages(name string) map[string]string {
	return MapErrorRequest{
		"name": {
			"required": "Name field is required",
		},
		"description": {
			"required": "Description field is required",
		},
		"address": {
			"required": "Address field is required",
		},
	}[name]
}

func (RegisterStoreRequest) FieldName(name string) string {
	return map[string]string{
		"Name":        "name",
		"Description": "description",
		"Address":     "address",
	}[name]
}
