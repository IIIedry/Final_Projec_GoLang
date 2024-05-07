package handlers

import (
	"Application"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.services.Products.GetAll(c)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"products": products,
	},
	)
}

func (h *Handler) GetProductById(c *gin.Context) {
	id := c.Param("id")
	id_num, err := strconv.ParseInt(id, 10, 64)
	product, err := h.services.Products.GetById(int(id_num), c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"product": product,
	},
	)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var product Application.Product
	if err := c.ShouldBind(&product); err != nil {
		log.Println(err)
	}
	id := c.Param("id")
	id_num, err := strconv.ParseInt(id, 10, 64)
	product.ID = int(id_num)
	updated, err := h.services.Products.Update(product, c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"updated": updated,
	},
	)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	id_num, err := strconv.ParseInt(id, 10, 64)
	deleted, err := h.services.Products.Delete(int(id_num), c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"deleted": deleted,
	},
	)
}
