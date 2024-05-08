package repository

import (
	"Application"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
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
}

type Products interface {
	Create(product Application.Product, ctx *gin.Context) (string, error)
	GetAll(ctx *gin.Context) ([]Application.Product, error)
	GetById(id int, ctx *gin.Context) (Application.Product, error)
	Update(product Application.Product, ctx *gin.Context) (int, error)
	Delete(id int, ctx *gin.Context) (bool, error)
}

type Repository struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Products:      NewProductsPg(db),
		Authorization: NewUsersPg(db),
		Administrator: NewAdminPg(db),
	}
}
