package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
}

type Administrator interface {
}

type Orders interface {
	Create(product Application.Order, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	GetById(id int, ctx *gin.Context) (Application.Order, error)
	Delete(id int, ctx *gin.Context) (bool, error)
	Update(order Application.Order, ctx *gin.Context) (int, error)
}

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	GetById(id int, ctx *gin.Context) (Application.Product, error)
	Delete(id int, ctx *gin.Context) (bool, error)
	Update(product Application.Product, ctx *gin.Context) (int, error)
}

type Service struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Products: NewProductService(repos.Products)}
}
