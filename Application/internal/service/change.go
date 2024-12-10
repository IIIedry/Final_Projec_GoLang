package service

import (
	"Application"
	"Application/internal/repository"
	"github.com/gin-gonic/gin"
)

type ChangeService struct {
	repo repository.ProductChange
}

func (s *ChangeService) GetChanges(i int, c *gin.Context) ([]Application.Change, error) {
	return s.repo.GetChanges(i, c)
}

func (s *ChangeService) Create(change Application.Change, ctx *gin.Context) (string, error) {
	return s.repo.Create(change, ctx)
}

func NewChangeService(repo repository.ProductChange) *ChangeService {
	return &ChangeService{repo: repo}
}
