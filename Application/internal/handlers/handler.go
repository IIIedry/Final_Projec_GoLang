package handlers

import (
	"Application/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	home := router.Group("/")
	{
		home.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}
	auth := router.Group("/auth")
	{
		auth.POST("/create", h.AddUser)
		auth.POST("/login", h.Login)
		auth.GET("/:id", h.GetUserById)
	}
	product := router.Group("/products")
	{
		product.POST("/create", h.AddProduct)
		product.GET("/", h.GetAllProducts)
		product.GET("/:id", h.GetProductById)
		product.PUT("/:id/update", h.UpdateProduct)
		product.DELETE("/:id/delete", h.DeleteProduct)
	}
	admin := router.Group("/admin")
	{
		admin.POST("/", h.Login_Admin)
		admin.GET("/users/", h.GetAllUser)
		admin.PUT("/users/:id/update_role", h.UpdateUserRole)
		admin.GET("/users/:id", h.GetUserById)
	}

	return router
}
