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
		auth.GET("/", h.GetAllUser)
		auth.POST("/login", h.Login)
	}
	product := router.Group("/products")
	{
		product.POST("/create", h.AddProduct)
		product.GET("/", h.GetAllProducts)
		product.GET("/:id", h.GetProductById)
		product.PUT("/:id/update", h.UpdateProduct)
		product.DELETE("/:id/delete", h.DeleteProduct)
	}

	return router
}
