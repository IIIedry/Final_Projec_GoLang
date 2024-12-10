package Application

import "github.com/jackc/pgtype"

type Product struct {
	ID          int         `json:"ID" db:"product_id"`
	Name        string      `json:"Name" db:"name"`
	Description string      `json:"Description" db:"description"`
	Price       string      `json:"Price" db:"price"`
	Count       string      `json:"Count" db:"count"`
	CreatedAt   pgtype.Time `json:"CreatedAt" db:"created_at"`
}
