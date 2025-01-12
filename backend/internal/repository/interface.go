package repository

type IRepository[T any] interface {
	GetAll() ([]*T, error)
	GetByID(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(uint) error
}

type IOrderRepository[T any] interface {
	IRepository[T]
	GetActiveOrders() ([]*T, error)
	GetCompletedOrders() ([]*T, error)
}
