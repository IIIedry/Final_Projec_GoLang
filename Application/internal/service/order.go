package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

//import (
//	"Application"
//	"Application/internal/repository"
//	"github.com/gin-gonic/gin"
//)

type OrderService struct {
	repo repository.Orders
}

func (s *OrderService) GetById(id int, ctx *gin.Context) (Application.Order, error) {
	return s.repo.GetById(id, ctx)
}

func NewOrderService(repo repository.Orders) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order Application.Order, ctx *gin.Context) (string, error) {
	return s.repo.Create(order, ctx)
}

func (s *OrderService) GetAll(ctx *gin.Context) ([]Application.Order, error) {
	return s.repo.GetAll(ctx)
}
func (s *OrderService) Update(order Application.Order, ctx *gin.Context) (int, error) {
	return s.repo.Update(order, ctx)
}
func (s *OrderService) Delete(id int, ctx *gin.Context) (bool, error) {
	return s.repo.Delete(id, ctx)
}
