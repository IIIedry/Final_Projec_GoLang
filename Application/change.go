package Application

import "github.com/jackc/pgtype"

type Change struct {
	ID        string      `json:"id" db:"product_change_id"`
	ProductID string      `json:"product_id" db:"product_id"`
	Change    string      `json:"change" db:"change"`
	UpdateAt  pgtype.Time `json:"updated_at" db:"updated_at"`
}
