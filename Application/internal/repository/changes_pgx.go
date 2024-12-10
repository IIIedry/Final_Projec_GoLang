package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type ChangesPgx struct {
	conn *pgx.Conn
}

func (r *ChangesPgx) GetChanges(i int, c *gin.Context) ([]Application.Change, error) {
	changes := make([]Application.Change, 0)
	rows, err := r.conn.Query(c, "SELECT * FROM product_changes WHERE product_id = $1", i)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var change Application.Change
		err := rows.Scan(&change.ID, &change.ProductID, &change.Change, &change.UpdateAt)
		changes = append(changes, change)
		if err != nil {
			return nil, err
		}
	}
	return changes, nil
}

func NewChangesPg(conn *pgx.Conn) *ChangesPgx {
	return &ChangesPgx{conn: conn}
}

func (r *ChangesPgx) Create(changes Application.Change, ctx *gin.Context) (string, error) {
	row := r.conn.QueryRow(ctx, "INSERT INTO product_changes (product_id, change) VALUES ($1, $2) RETURNING change", changes.ProductID, changes.Change)
	var change string
	if err := row.Scan(&change); err != nil {
		return "0", err
	}
	return "Success", nil
}
