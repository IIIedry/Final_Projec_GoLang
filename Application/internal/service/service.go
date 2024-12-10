package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	GetById(id int, ctx *gin.Context) (Application.Product, error)
	Delete(id int, ctx *gin.Context) (bool, error)
	Update(product Application.Product, ctx *gin.Context) (int, map[string]interface{}, error)
}

type ProductChange interface {
	Create(change Application.Change, ctx *gin.Context) (string, error)
	GetChanges(i int, c *gin.Context) ([]Application.Change, error)
}

type Service struct {
	Products
	ProductChange
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Products:      NewProductService(repos.Products),
		ProductChange: NewChangeService(repos.ProductChange),
	}
}
