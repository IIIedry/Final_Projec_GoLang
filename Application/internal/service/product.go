package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repo repository.Products
}

func NewProductService(repo repository.Products) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product Application.Product, ctx *gin.Context) (string, error) {
	return s.repo.Create(product, ctx)
}
