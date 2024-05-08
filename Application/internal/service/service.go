package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user Application.User, ctx *gin.Context) (string, error)
	AuthenticateUser(username, password string, ctx *gin.Context) (*Application.User, error)
	GetUserById(id int, ctx *gin.Context) (Application.User, error)
}

type Administrator interface {
	GetAllUser(ctx *gin.Context) ([]Application.User, error)
	UpdateUserRole(userID int, newRole string, ctx *gin.Context) error
	IsAdmin(username, password string, ctx *gin.Context) (bool, error)
}

type Orders interface {
	Create(order Application.Order, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Order, error)
	GetById(id int, ctx *gin.Context) (Application.Order, error)
	Delete(id int, ctx *gin.Context) (bool, error)
	Update(product Application.Order, ctx *gin.Context) (int, error)
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
		Products:      NewProductService(repos.Products),
		Authorization: NewUserService(repos.Authorization),
		Administrator: NewAdminService(repos.Administrator),
		Orders:        NewOrderService(repos.Orders),
	}
}
