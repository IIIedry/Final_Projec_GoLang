package repository

import "github.com/jackc/pgx/v4"

type Authorization interface {
}

type Administrator interface {
}

type Orders interface {
}

type Products interface{}

type Repository struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{}
}
