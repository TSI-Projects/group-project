package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TSI-Projects/group-project/internal/api/handlers"
	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/TSI-Projects/group-project/pkg/validation"
	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type mockOrderTypeRepo struct {
	getAllFunc  func() ([]*repository.OrderType, error)
	createFunc  func(*repository.OrderType) error
	deleteFunc  func(uint) error
	getByIDFunc func(uint) (*repository.OrderType, error)
	updateFunc  func(*repository.OrderType) error
}

func (m *mockOrderTypeRepo) GetAll() ([]*repository.OrderType, error) {
	if m.getAllFunc != nil {
		return m.getAllFunc()
	}
	return nil, nil
}

func (m *mockOrderTypeRepo) GetByID(id uint) (*repository.OrderType, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(id)
	}
	return nil, nil
}

func (m *mockOrderTypeRepo) Create(orderType *repository.OrderType) error {
	if m.createFunc != nil {
		return m.createFunc(orderType)
	}
	return nil
}

func (m *mockOrderTypeRepo) Update(orderType *repository.OrderType) error {
	if m.updateFunc != nil {
		return m.updateFunc(orderType)
	}
	return nil
}

func (m *mockOrderTypeRepo) Delete(id uint) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id)
	}
	return nil
}

func TestGetOrderTypes_Success(t *testing.T) {
	mockRepo := &mockOrderTypeRepo{
		getAllFunc: func() ([]*repository.OrderType, error) {
			return []*repository.OrderType{
				{ID: 1, FullName: "Type1"},
				{ID: 2, FullName: "Type2"},
			}, nil
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/order-types", nil)
	rr := httptest.NewRecorder()

	handler.GetOrderTypes(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "expected status 200 OK")

	var respData map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &respData)
	assert.NoError(t, err, "failed to unmarshal response")

	orderTypes, exists := respData["order_types"].([]interface{})
	assert.True(t, exists, "key 'order_types' doesn't exist in response")
	assert.Len(t, orderTypes, 2, "expected 2 elements in 'order_types'")
}

func TestGetOrderTypes_Failure_GetAllError(t *testing.T) {
	mockRepo := &mockOrderTypeRepo{
		getAllFunc: func() ([]*repository.OrderType, error) {
			return nil, errors.New("database error")
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/order-types", nil)
	rr := httptest.NewRecorder()

	handler.GetOrderTypes(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "expected status 500 Internal Server Error")
}

// ----- CreateOrderType Tests -----

func TestCreateOrderType_Success(t *testing.T) {
	mockRepo := &mockOrderTypeRepo{
		createFunc: func(ot *repository.OrderType) error {
			ot.ID = 1
			return nil
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	orderType := &repository.OrderType{FullName: "NewType"}
	bodyBytes, _ := json.Marshal(orderType)
	req := httptest.NewRequest(http.MethodPost, "/order-types", bytes.NewReader(bodyBytes))
	rr := httptest.NewRecorder()

	handler.CreateOrderType(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateOrderType_Failure_DecodeError(t *testing.T) {
	handler := &handlers.Handler{
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodPost, "/order-types", bytes.NewReader([]byte("{invalid-json")))
	rr := httptest.NewRecorder()

	handler.CreateOrderType(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

// ----- DeleteOrderType Tests -----

func TestDeleteOrderType_Success(t *testing.T) {
	mockRepo := &mockOrderTypeRepo{
		deleteFunc: func(id uint) error {
			return nil
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodDelete, "/order-types/1", nil)
	// Устанавливаем mux.Vars вручную для теста
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	rr := httptest.NewRecorder()
	handler.DeleteOrderType(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteOrderType_Failure_EmptyID(t *testing.T) {
	handler := &handlers.Handler{
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodDelete, "/order-types/", nil)
	req = mux.SetURLVars(req, map[string]string{"id": ""})

	rr := httptest.NewRecorder()
	handler.DeleteOrderType(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

// ----- GetOrderTypeByID Tests -----

func TestGetOrderTypeByID_Success(t *testing.T) {
	expectedType := &repository.OrderType{ID: 1, FullName: "Type1"}
	mockRepo := &mockOrderTypeRepo{
		getByIDFunc: func(id uint) (*repository.OrderType, error) {
			if id == 1 {
				return expectedType, nil
			}
			return nil, errors.New("not found")
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/order-types/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	handler.GetOrderTypeByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var respData map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &respData)
	assert.NoError(t, err)

	orderTypeData, exists := respData["order_type"].(map[string]interface{})
	assert.True(t, exists)

	assert.Equal(t, float64(expectedType.ID), orderTypeData["id"])
	assert.Equal(t, expectedType.FullName, orderTypeData["full_name"])
}

func TestGetOrderTypeByID_Failure_InvalidID(t *testing.T) {
	handler := &handlers.Handler{
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodGet, "/order-types/invalid", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "invalid"})
	rr := httptest.NewRecorder()

	handler.GetOrderTypeByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// ----- UpdateOrderType Tests -----

func TestUpdateOrderType_Success(t *testing.T) {
	mockRepo := &mockOrderTypeRepo{
		updateFunc: func(ot *repository.OrderType) error {
			return nil
		},
	}

	handler := &handlers.Handler{
		OrderTypeRepo: mockRepo,
		Validator:     validation.NewValidatorClient(),
	}

	orderType := &repository.OrderType{ID: 1, FullName: "UpdatedType"}
	bodyBytes, _ := json.Marshal(orderType)
	req := httptest.NewRequest(http.MethodPut, "/order-types", bytes.NewReader(bodyBytes))
	rr := httptest.NewRecorder()

	handler.UpdateOrderType(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateOrderType_Failure_DecodeError(t *testing.T) {
	handler := &handlers.Handler{
		Validator: validation.NewValidatorClient(),
	}

	req := httptest.NewRequest(http.MethodPut, "/order-types", bytes.NewReader([]byte("{invalid-json")))
	rr := httptest.NewRecorder()

	handler.UpdateOrderType(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
