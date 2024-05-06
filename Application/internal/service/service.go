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
}

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	//GetById(id int) (Application.Product, error)
	//Delete(id int) (bool, error)
	//Update(product Application.Product) (int, error)
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
