package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) Login_Admin(c *gin.Context) {
	var req struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	isAdmin, err := h.services.Administrator.IsAdmin(req.Username, req.Password, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if !isAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"is_admin": isAdmin})
}
func (h *Handler) UpdateUserRole(c *gin.Context) {
	var req struct {
		UserID  int    `json:"ID"`
		NewRole string `json:"Role"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.services.Administrator.UpdateUserRole(req.UserID, req.NewRole, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user role updated successfully"})
}
func (h *Handler) GetAllUser(c *gin.Context) {
	users, err := h.services.Administrator.GetAllUser(c)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{
		"user": users,
	},
	)
}
