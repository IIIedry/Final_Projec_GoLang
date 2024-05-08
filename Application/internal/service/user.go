package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.Authorization
}

func NewUserService(repo repository.Authorization) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user Application.User, ctx *gin.Context) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
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

func (s *UserService) IsAdmin(username, password string, ctx *gin.Context) (bool, error) {
	isAdmin, err := s.repo.IsAdmin(username, password, ctx)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}

func (s *UserService) UpdateUserRole(userID int, newRole string, ctx *gin.Context) error {
	err := s.repo.UpdateUserRole(userID, newRole, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserById(id int, ctx *gin.Context) (Application.User, error) {
	return s.repo.GetUserById(id, ctx)
}
