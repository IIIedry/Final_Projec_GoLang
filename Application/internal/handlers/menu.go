package handlers

import (
	"Application"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) AddProduct(c *gin.Context) {
	var product Application.Product
	if err := c.ShouldBind(&product); err != nil {
		log.Println(err)
	}
	name, err := h.services.Products.Create(product, c)
	if err != nil {
		log.Println(err)
	}
	c.String(200, "Product created successfully")
	c.JSON(200, gin.H{
		"name": name,
	},
	)
}
