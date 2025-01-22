package fakeModels

import "github.com/TSI-Projects/group-project/internal/repository"

type MockCustomerRepo struct{}

func (m *MockCustomerRepo) GetAll() ([]*repository.Customer, error) {
	return nil, nil
}
func (m *MockCustomerRepo) GetByID(id uint) (*repository.Customer, error) {
	return nil, nil
}
func (m *MockCustomerRepo) Create(c *repository.Customer) error {
	return nil
}
func (m *MockCustomerRepo) Update(c *repository.Customer) error {
	return nil
}
func (m *MockCustomerRepo) Delete(id uint) error {
	return nil
}

type MockOrderStatusRepo struct{}

func (m *MockOrderStatusRepo) GetAll() ([]*repository.OrderStatus, error) {
	return nil, nil
}
func (m *MockOrderStatusRepo) GetByID(id uint) (*repository.OrderStatus, error) {
	return nil, nil
}
func (m *MockOrderStatusRepo) Create(os *repository.OrderStatus) error {
	return nil
}
func (m *MockOrderStatusRepo) Update(os *repository.OrderStatus) error {
	return nil
}
func (m *MockOrderStatusRepo) Delete(id uint) error {
	return nil
}
