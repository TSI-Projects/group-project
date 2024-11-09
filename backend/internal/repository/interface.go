package repository

type IRepository[T any] interface {
	GetAll() ([]*T, error)
	GetByID(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(uint) error
}
