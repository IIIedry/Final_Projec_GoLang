package handlers

import (
	"Application"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.services.Authorization.AuthenticateUser(req.Username, req.Password, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	id_num, err := strconv.ParseInt(id, 10, 64)
	user, err := h.services.Authorization.GetUserById(int(id_num), c)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"user": user,
	},
	)
}
