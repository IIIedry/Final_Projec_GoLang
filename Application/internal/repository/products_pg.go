package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
)

type ProductsPgx struct {
	conn *pgx.Conn
}

func NewProductsPg(conn *pgx.Conn) *ProductsPgx {
	return &ProductsPgx{conn: conn}
}

func (r *ProductsPgx) Create(product Application.Product, ctx *gin.Context) (string, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return "0", err
	}
	defer tx.Rollback(ctx)

	var name string
	row := tx.QueryRow(ctx, "INSERT INTO products (name, description, price, count) VALUES ($1, $2, $3, $4) RETURNING name", product.Name, product.Description, product.Price, product.Count)
	if err = row.Scan(&name); err != nil {
		tx.Rollback(ctx)
		return "0", err
	}
	if err != nil {
		return "0", err
	}
	return name, tx.Commit(ctx)
}

func (r *ProductsPgx) GetAll(ctx *gin.Context) ([]Application.Product, error) {
	var products []Application.Product
	tx, err := r.conn.Begin(ctx)
	rows, err := tx.Query(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Application.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.CreatedAt)
		log.Println(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
		log.Println(products)
	}
	tx.Commit(ctx)
	return products, nil
}
