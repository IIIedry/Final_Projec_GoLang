package Application

import "github.com/jackc/pgtype"

type Order struct {
	ID        int         `json:"ID" db:"id"`
	UserID    int         `json:"UserID" db:"user_id"`
	ProductID int         `json:"ProductID" db:"product_id"`
	Total     int         `json:"Total" db:"total"`
	Status    string      `json:"Status" db:"status"`
	CreatedAt pgtype.Time `json:"Created_at" db:"created_at"`
}
