package repository

type IRepository[T any] interface {
	GetAll() ([]*T, error)
	GetByID(int) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(int) error
}
