package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type OrderService struct {
	repo repository.Order 
}

func (s *OrderService) GetById(id int, ctx *gin.Context) (Application.Product, error) {
	return s.repo.GetById(id, ctx)
}

func NewProductService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(product Application.Order, ctx *gin.Context) (string, error) {
	return s.repo.Create(product, ctx)
}

func (s *OrderService) GetAll(ctx *gin.Context) ([]Application.Order, error) {
	return s.repo.GetAll(ctx)
}

func (s *OrderService) Update(product Application.Order, ctx *gin.Context) (int, error) {
	return s.repo.Update(product, ctx)
}

func (s *OrderService) Delete(id int, ctx *gin.Context) (bool, error) {
	return s.repo.Delete(id, ctx)
}
