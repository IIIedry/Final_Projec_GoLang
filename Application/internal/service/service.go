package service

import "Application/internal/repository"

type Authorization interface {
}

type Administrator interface {
}

type Orders interface {
}

type Products interface{}

type Service struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
