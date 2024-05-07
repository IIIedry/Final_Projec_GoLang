package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProductService struct {
	repo repository.Products
}

func (s *ProductService) GetById(id int, ctx *gin.Context) (Application.Product, error) {
	return s.repo.GetById(id, ctx)
}

func NewProductService(repo repository.Products) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product Application.Product, ctx *gin.Context) (string, error) {
	return s.repo.Create(product, ctx)
}

func (s *ProductService) GetAll(ctx *gin.Context) ([]Application.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) Update(product Application.Product, ctx *gin.Context) (int, error) {
	return s.repo.Update(product, ctx)
}

func (s *ProductService) Delete(id int, ctx *gin.Context) (bool, error) {
	return s.repo.Delete(id, ctx)
}
