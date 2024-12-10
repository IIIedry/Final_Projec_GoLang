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
	changes, err := h.services.ProductChange.GetChanges(int(id_num), c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"id":      id,
		"product": product,
		"changes": changes,
	},
	)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var product Application.Product
	if err := c.ShouldBind(&product); err != nil {
		log.Println(err)
	}
	_, result, err := h.services.Products.Update(product, c)
	var change string
	for key, value := range result {
		if err, ok := value.(string); !ok {
			log.Println(err)
			change += key + ": " + strconv.Itoa(value.(int)) + ", "
		} else {
			change += key + ": " + value.(string) + ", "
		}

	}
	id := c.Param("id")
	history, err := h.services.ProductChange.Create(Application.Change{
		ProductID: id,
		Change:    change,
	}, c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"status": history,
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
