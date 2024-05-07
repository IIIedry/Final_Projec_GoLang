package repository

import (
	"Application"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"strconv"
	"strings"
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

func (r *ProductsPgx) GetById(id int, ctx *gin.Context) (Application.Product, error) {
	var product Application.Product
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return product, err
	}
	row := tx.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", id)
	err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count, &product.CreatedAt)
	if err != nil {
		return product, err
	}
	tx.Commit(ctx)
	return product, nil
}

func (r *ProductsPgx) Update(product Application.Product, ctx *gin.Context) (int, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return 0, err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if product.Name != "" {
		setValues = append(setValues, "name = $"+strconv.Itoa(argId))
		args = append(args, product.Name)
		argId++
	}
	if product.Description != "" {
		setValues = append(setValues, "description = $"+strconv.Itoa(argId))
		args = append(args, product.Description)
		argId++
	}
	if product.Price != 0 {
		setValues = append(setValues, "price = $"+strconv.Itoa(argId))
		args = append(args, product.Price)
		argId++
	}
	if product.Count != 0 {
		setValues = append(setValues, "count = $"+strconv.Itoa(argId))
		args = append(args, product.Count)
		argId++
	}
	args = append(args, product.ID)
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d", setQuery, argId)
	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return 0, nil
	}

	tx.Commit(ctx)
	return 1, nil
}

func (r *ProductsPgx) Delete(id int, ctx *gin.Context) (bool, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	tx.Commit(ctx)
	return true, nil
}
