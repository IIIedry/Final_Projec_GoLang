package repository

type Authorization interface {
}

type Administrator interface {
}

type Orders interface {
}

type Products interface{}

type Repository struct {
	Authorization
	Administrator
	Orders
	Products
}

func NewRepository() *Repository {
	return &Repository{}
}
