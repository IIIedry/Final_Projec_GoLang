package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo repository.Authorization
}

func NewUserService(repo repository.Authorization) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user Application.User, ctx *gin.Context) (string, error) {
	return s.repo.CreateUser(user, ctx)
}

func (s *UserService) GetAllUser(ctx *gin.Context) ([]Application.User, error) {
	return s.repo.GetAllUser(ctx)
}

func (s *UserService) AuthenticateUser(username, password string, ctx *gin.Context) (*Application.User, error) {
	user, err := s.repo.AuthenticateUser(username, password, ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
