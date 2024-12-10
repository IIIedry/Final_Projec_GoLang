package repository

import (
	"Application"
	"context"
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
	rows, err := r.conn.Query(ctx, "SELECT product_id, name, description, price, count, created_at FROM products")
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
	return products, nil
}

func (r *ProductsPgx) GetById(id int, ctx *gin.Context) (Application.Product, error) {
	var product Application.Product
	product, err := r.findProduct(int64(id))
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductsPgx) findProduct(id int64) (Application.Product, error) {
	var product Application.Product
	row := r.conn.QueryRow(context.Background(), "SELECT product_id, name, description, price, count FROM products WHERE product_id = $1", id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Count)
	log.Println(product.ID, product.Name, product.Description, product.Price, product.Count)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductsPgx) Update(product Application.Product, ctx *gin.Context) (int, map[string]interface{}, error) {
	tx, err := r.conn.Begin(ctx)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	result := make(map[string]interface{})
	argId := 1
	id := ctx.Param("id")
	id_num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic("err")
	}

	if product.Name != "" {
		setValues = append(setValues, "name = $"+strconv.Itoa(argId))
		args = append(args, product.Name)
		result["name"] = product.Name
		argId++
	}
	if product.Description != "" {
		setValues = append(setValues, "description = $"+strconv.Itoa(argId))
		args = append(args, product.Description)
		result["description"] = product.Description
		argId++
	}
	if product.Price != "" {
		setValues = append(setValues, "price = $"+strconv.Itoa(argId))
		args = append(args, product.Price)
		result["price"] = product.Price
		argId++
	}
	if product.Count != "" {
		setValues = append(setValues, "count = $"+strconv.Itoa(argId))
		result["count"] = product.Count
		args = append(args, product.Count)
		argId++
	}
	args = append(args, id_num)
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE products SET %s WHERE product_id = $%d", setQuery, argId)
	_, err = tx.Exec(ctx, query, args...)

	if err != nil {
		return 0, nil, err
	}

	log.Println(id_num, result)
	err = tx.Commit(ctx)
	if err != nil {
		return 0, nil, err
	}
	return int(id_num), result, nil
}

func (r *ProductsPgx) Delete(id int, ctx *gin.Context) (bool, error) {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(ctx, "DELETE FROM products WHERE product_id = $1", id)
	if err != nil {
		return false, err
	}
	tx.Commit(ctx)
	return true, nil
}
