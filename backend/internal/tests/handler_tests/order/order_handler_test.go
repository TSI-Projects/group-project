package handlers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TSI-Projects/group-project/internal/api/handlers"
	"github.com/TSI-Projects/group-project/internal/repository"
	fakeModels "github.com/TSI-Projects/group-project/internal/tests/handler_tests/fake/models"
	"github.com/TSI-Projects/group-project/pkg/validation"
	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//	Мок-репозиторий
//
// -------------------------
type mockOrderRepo struct {
	getAllFunc             func() ([]*repository.Order, error)
	getActiveOrdersFunc    func() ([]*repository.Order, error)
	getCompletedOrdersFunc func() ([]*repository.Order, error)
	getByIDFunc            func(uint) (*repository.Order, error)
	createFunc             func(*repository.Order) error
	updateFunc             func(*repository.Order) error
	deleteFunc             func(uint) error
}

func (m *mockOrderRepo) GetAll() ([]*repository.Order, error) {
	if m.getAllFunc != nil {
		return m.getAllFunc()
	}
	return nil, nil
}
func (m *mockOrderRepo) GetActiveOrders() ([]*repository.Order, error) {
	if m.getActiveOrdersFunc != nil {
		return m.getActiveOrdersFunc()
	}
	return nil, nil
}
func (m *mockOrderRepo) GetCompletedOrders() ([]*repository.Order, error) {
	if m.getCompletedOrdersFunc != nil {
		return m.getCompletedOrdersFunc()
	}
	return nil, nil
}
func (m *mockOrderRepo) GetByID(id uint) (*repository.Order, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(id)
	}
	return nil, nil
}
func (m *mockOrderRepo) Create(o *repository.Order) error {
	if m.createFunc != nil {
		return m.createFunc(o)
	}
	return nil
}
func (m *mockOrderRepo) Update(o *repository.Order) error {
	if m.updateFunc != nil {
		return m.updateFunc(o)
	}
	return nil
}
func (m *mockOrderRepo) Delete(id uint) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id)
	}
	return nil
}

// -------------------------
//
//	GetOrders
//
// -------------------------
func TestGetOrders_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getAllFunc: func() ([]*repository.Order, error) {
			return []*repository.Order{
				{
					ID:          1,
					OrderTypeID: 1,
					WorkerID:    2,
					Reason:      "Some reason",
					Defect:      "Some defect",
					ItemName:    "Item1",
					TotalPrice:  100.0,
					Prepayment:  50.0,
				},
				{
					ID:          2,
					OrderTypeID: 1,
					WorkerID:    2,
					Reason:      "Another reason",
					Defect:      "Another defect",
					ItemName:    "Item2",
					TotalPrice:  200.0,
					Prepayment:  70.0,
				},
			}, nil
		},
	}

	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/api/orders", nil)
	rr := httptest.NewRecorder()

	h.GetOrders(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)

	orders, ok := resp["orders"].([]interface{})
	assert.True(t, ok, "expected 'orders' in response")
	assert.Len(t, orders, 2, "expected 2 orders")
}

func TestGetOrders_Failure_DBError(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getAllFunc: func() ([]*repository.Order, error) {
			return nil, errors.New("DB error")
		},
	}
	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/api/orders", nil)
	rr := httptest.NewRecorder()

	h.GetOrders(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// -------------------------
//
//	GetActiveOrders
//
// -------------------------
func TestGetActiveOrders_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getActiveOrdersFunc: func() ([]*repository.Order, error) {
			return []*repository.Order{
				{
					ID:          10,
					OrderTypeID: 1,
					WorkerID:    2,
					Reason:      "Active reason",
					Defect:      "Active defect",
					ItemName:    "Active item",
					TotalPrice:  300.0,
					Prepayment:  100.0,
				},
			}, nil
		},
	}

	h := &handlers.Handler{OrderRepo: mockRepo, Validator: validation.NewValidatorClient()}

	req := httptest.NewRequest(http.MethodGet, "/api/orders/active", nil)
	rr := httptest.NewRecorder()

	h.GetActiveOrders(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)

	orders, ok := resp["orders"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, orders, 1)
}

func TestGetActiveOrders_Failure_DBError(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getActiveOrdersFunc: func() ([]*repository.Order, error) {
			return nil, errors.New("DB error")
		},
	}

	h := &handlers.Handler{OrderRepo: mockRepo}
	req := httptest.NewRequest(http.MethodGet, "/api/orders/active", nil)
	rr := httptest.NewRecorder()

	h.GetActiveOrders(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// -------------------------
//
//	GetCompletedOrders
//
// -------------------------
func TestGetCompletedOrders_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getCompletedOrdersFunc: func() ([]*repository.Order, error) {
			return []*repository.Order{
				{
					ID:          5,
					OrderTypeID: 1,
					WorkerID:    2,
					Reason:      "Completed reason",
					Defect:      "Completed defect",
					ItemName:    "Completed item",
					TotalPrice:  999.0,
					Prepayment:  500.0,
				},
			}, nil
		},
	}

	h := &handlers.Handler{OrderRepo: mockRepo, Validator: validation.NewValidatorClient()}

	req := httptest.NewRequest(http.MethodGet, "/api/orders/completed", nil)
	rr := httptest.NewRecorder()

	h.GetCompletedOrders(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)

	orders, ok := resp["orders"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, orders, 1)
}

func TestGetCompletedOrders_Failure_DBError(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getCompletedOrdersFunc: func() ([]*repository.Order, error) {
			return nil, errors.New("DB error")
		},
	}
	h := &handlers.Handler{OrderRepo: mockRepo}

	req := httptest.NewRequest(http.MethodGet, "/api/orders/completed", nil)
	rr := httptest.NewRecorder()

	h.GetCompletedOrders(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// -------------------------
//
//	CreateOrder
//
// -------------------------
func TestCreateOrder_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		createFunc: func(o *repository.Order) error {
			o.ID = 1
			return nil
		},
	}
	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	// Заполняем все обязательные поля (validate:"required")
	newOrder := &repository.Order{
		OrderTypeID: 1,
		WorkerID:    2,
		Reason:      "Reason for repair",
		Defect:      "Does not turn on",
		ItemName:    "Laptop",
		TotalPrice:  500.00,
		Prepayment:  100.00,
		// Опциональные поля:
		OrderStatusID: 1,            // validate:"omitempty"
		CustomerID:    101,          // validate:"omitempty"
		CreatedAt:     &time.Time{}, // example
	}
	body, _ := json.Marshal(newOrder)

	req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.CreateOrder(rr, req)

	// При успешном создании 200 OK (в вашем коде)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateOrder_Failure_DecodeError(t *testing.T) {
	mockRepo := &mockOrderRepo{}
	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewReader([]byte("{bad json")))
	rr := httptest.NewRecorder()

	h.CreateOrder(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestCreateOrder_Failure_ValidationError(t *testing.T) {
	// Укажем не все обязательные поля
	invalidOrder := &repository.Order{
		// workerID and reason, defect, itemName, totalPrice, prepayment — обязательны.
		// Допустим, не укажем WorkerID
		OrderTypeID: 1,
		Reason:      "No worker",
		Defect:      "Some defect",
		ItemName:    "ItemName",
		TotalPrice:  100.0,
		Prepayment:  50.0,
	}

	mockRepo := &mockOrderRepo{
		createFunc: func(o *repository.Order) error {
			return nil // неважно, всё равно не дойдет из-за ошибки валидации
		},
	}
	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	body, _ := json.Marshal(invalidOrder)
	req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.CreateOrder(rr, req)

	// Судя по коду, при ошибке валидации приходит 400
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// -------------------------
//
//	DeleteOrder
//
// -------------------------
func TestDeleteOrder_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		deleteFunc: func(id uint) error {
			return nil
		},
	}
	h := &handlers.Handler{OrderRepo: mockRepo}

	req := httptest.NewRequest(http.MethodDelete, "/api/order/10", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "10"})

	rr := httptest.NewRecorder()
	h.DeleteOrder(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteOrder_Failure_EmptyID(t *testing.T) {
	h := &handlers.Handler{}

	req := httptest.NewRequest(http.MethodDelete, "/api/order/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": ""})

	rr := httptest.NewRecorder()
	h.DeleteOrder(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDeleteOrder_Failure_InvalidID(t *testing.T) {
	h := &handlers.Handler{}

	req := httptest.NewRequest(http.MethodDelete, "/api/order/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	rr := httptest.NewRecorder()
	h.DeleteOrder(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestDeleteOrder_Failure_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{
		deleteFunc: func(id uint) error {
			return errors.New("repo delete error")
		},
	}
	h := &handlers.Handler{OrderRepo: mockRepo}

	req := httptest.NewRequest(http.MethodDelete, "/api/order/5", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "5"})

	rr := httptest.NewRecorder()
	h.DeleteOrder(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// -------------------------
//
//	GetOrderByID
//
// -------------------------
func TestGetOrderByID_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getByIDFunc: func(id uint) (*repository.Order, error) {
			return &repository.Order{
				ID:          id,
				OrderTypeID: 1,
				WorkerID:    2,
				Reason:      "Reason",
				Defect:      "Defect",
				ItemName:    "ItemName",
				TotalPrice:  1234.56,
				Prepayment:  500.00,
			}, nil
		},
	}
	h := &handlers.Handler{OrderRepo: mockRepo}

	req := httptest.NewRequest(http.MethodGet, "/api/order/7", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "7"})

	rr := httptest.NewRecorder()
	h.GetOrderByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)

	orderData, ok := resp["order"].(map[string]interface{})
	assert.True(t, ok, "expected 'order' key in response")
	assert.Equal(t, float64(7), orderData["id"])
	assert.Equal(t, "Reason", orderData["reason"])
	assert.Equal(t, "Defect", orderData["defect"])
	assert.Equal(t, "ItemName", orderData["item_name"])
	assert.Equal(t, 1234.56, orderData["total_price"])
	assert.Equal(t, 500.00, orderData["prepayment"])
}

func TestGetOrderByID_Failure_EmptyID(t *testing.T) {
	h := &handlers.Handler{}

	req := httptest.NewRequest(http.MethodGet, "/api/order/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": ""})

	rr := httptest.NewRecorder()
	h.GetOrderByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetOrderByID_Failure_InvalidID(t *testing.T) {
	h := &handlers.Handler{}

	req := httptest.NewRequest(http.MethodGet, "/api/order/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})

	rr := httptest.NewRecorder()
	h.GetOrderByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetOrderByID_Failure_NotFoundInRepo(t *testing.T) {
	mockRepo := &mockOrderRepo{
		getByIDFunc: func(id uint) (*repository.Order, error) {
			return nil, errors.New("not found")
		},
	}
	h := &handlers.Handler{OrderRepo: mockRepo}

	req := httptest.NewRequest(http.MethodGet, "/api/order/99", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "99"})

	rr := httptest.NewRecorder()
	h.GetOrderByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// -------------------------
//
//	UpdateOrder
//
// -------------------------
func TestUpdateOrder_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		updateFunc: func(o *repository.Order) error {
			return nil
		},
	}
	// В вашем коде UpdateOrder также вызывает OrderStatusRepo.Update(order.Status),
	// CustomerRepo.Update(order.Customer), и т.д. Их тоже можно замокать,
	// но для простоты тут оставим только OrderRepo.

	h := &handlers.Handler{
		OrderRepo:       mockRepo,
		OrderStatusRepo: &fakeModels.MockOrderStatusRepo{},
		CustomerRepo:    &fakeModels.MockCustomerRepo{},
		Validator:       validation.NewValidatorClient(),
	}

	updatedOrder := &repository.Order{
		ID:          10,
		OrderTypeID: 1,
		WorkerID:    2,
		Reason:      "Updated reason",
		Defect:      "Updated defect",
		ItemName:    "Updated item",
		TotalPrice:  200.0,
		Prepayment:  50.0,
	}
	body, _ := json.Marshal(updatedOrder)

	req := httptest.NewRequest(http.MethodPut, "/api/orders", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.UpdateOrder(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateOrder_Failure_DecodeError(t *testing.T) {
	h := &handlers.Handler{}

	req := httptest.NewRequest(http.MethodPut, "/api/orders", bytes.NewReader([]byte("{bad json")))
	rr := httptest.NewRecorder()

	h.UpdateOrder(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestUpdateOrder_Failure_ValidationError(t *testing.T) {
	invalidOrder := &repository.Order{
		OrderTypeID: 1,
		// WorkerID не указан
		Reason:     "Reason",
		Defect:     "Defect",
		ItemName:   "ItemName",
		TotalPrice: 100.0,
		Prepayment: 50.0,
	}

	mockRepo := &mockOrderRepo{}
	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	body, _ := json.Marshal(invalidOrder)
	req := httptest.NewRequest(http.MethodPut, "/api/orders", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.UpdateOrder(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateOrder_Failure_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{
		updateFunc: func(o *repository.Order) error {
			return errors.New("DB update error")
		},
	}

	h := &handlers.Handler{
		OrderRepo: mockRepo,
		Validator: validation.NewValidatorClient(),
	}

	order := &repository.Order{
		ID:          10,
		OrderTypeID: 1,
		WorkerID:    2,
		Reason:      "Reason",
		Defect:      "Defect",
		ItemName:    "ItemName",
		TotalPrice:  200.0,
		Prepayment:  50.0,
	}

	body, _ := json.Marshal(order)
	req := httptest.NewRequest(http.MethodPut, "/api/orders", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	h.UpdateOrder(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
