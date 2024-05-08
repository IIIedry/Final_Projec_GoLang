package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
	repo repository.Administrator
}

func NewAdminService(repo repository.Administrator) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) UpdateUserRole(userID int, newRole string, ctx *gin.Context) error {
	err := s.repo.UpdateUserRole(userID, newRole, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) GetAllUser(ctx *gin.Context) ([]Application.User, error) {
	return s.repo.GetAllUser(ctx)
}

func (s *AdminService) IsAdmin(username, password string, ctx *gin.Context) (bool, error) {
	isAdmin, err := s.repo.IsAdmin(username, password, ctx)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}
