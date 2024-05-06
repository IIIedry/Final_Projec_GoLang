package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Authorization interface {
}

type Administrator interface {
}

type Orders interface {
}

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
}

type Repository struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Products: NewProductsPg(db),
	}
}
