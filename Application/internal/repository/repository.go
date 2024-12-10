package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	GetById(id int, ctx *gin.Context) (Application.Product, error)
	Update(product Application.Product, ctx *gin.Context) (int, map[string]interface{}, error)
	Delete(id int, ctx *gin.Context) (bool, error)
}

type ProductChange interface {
	Create(change Application.Change, ctx *gin.Context) (string, error)
	GetChanges(i int, c *gin.Context) ([]Application.Change, error)
}

type Repository struct {
	Products
	ProductChange
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Products:      NewProductsPg(db),
		ProductChange: NewChangesPg(db),
	}
}
