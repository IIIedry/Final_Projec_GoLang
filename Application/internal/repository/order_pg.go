package repository

import (
	"Application"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
)

type OrdersPgx struct {
	conn *pgx.Conn
}

func NewOrdersPg(conn *pgx.Conn) *OrdersPgx {
	return &OrdersPgx{conn: conn}
}

func (r *OrdersPgx) Create(order Application.Order, ctx *gin.Context) (string, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "0", err
	}
	defer tx.Rollback(ctx)

	var name string
	row := tx.QueryRow(ctx, "INSERT INTO orders (product_id, user_id) values ($1, $2) returning id", order.ProductID, order.UserID)
	if err = row.Scan(&name); err != nil {
		tx.Rollback(ctx)
		return "0", err
	}
	if err != nil {
		return "0", err
	}
	return name, tx.Commit(ctx)
}

func (r *OrdersPgx) GetAll(ctx *gin.Context) ([]Application.Order, error) {
	var orders []Application.Order
	tx, err := r.conn.Begin(ctx)
	rows, err := tx.Query(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order Application.Order
		err = rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.CreatedAt, &order.Total)
		log.Println(order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
		log.Println(orders)
	}
	tx.Commit(ctx)
	return orders, nil
}

func (r *OrdersPgx) GetById(id int, ctx *gin.Context) (Application.Order, error) {
	var order Application.Order
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return order, err
	}
	row := tx.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", id)
	err = row.Scan(&order.ID, &order.ProductID, &order.UserID, &order.Total, &order.CreatedAt)
	if err != nil {
		return order, err
	}
	tx.Commit(ctx)
	return order, nil
}

func (r *OrdersPgx) Update(order Application.Order, ctx *gin.Context) (int, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("UPDATE products SET status = %s WHERE id = %d", order.Status, order.ID)
	_, err = tx.Exec(ctx, query)
	if err != nil {
		return 0, nil
	}
	tx.Commit(ctx)
	return 1, nil
}

func (r *OrdersPgx) Delete(id int, ctx *gin.Context) (bool, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(ctx, "DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	tx.Commit(ctx)
	return true, nil
}
