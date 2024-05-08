package handlers

import (
	"Application"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) AddUser(c *gin.Context) {
	var user Application.User
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err)
	}
	name, err := h.services.Authorization.CreateUser(user, c)
	if err != nil {
		log.Println(err)
	}
	c.String(200, "User created successfully")
	c.JSON(200, gin.H{
		"name": name,
	},
	)
}

func (h *Handler) GetAllUser(c *gin.Context) {
	users, err := h.services.Authorization.GetAllUser(c)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"user": users,
	},
	)
}
