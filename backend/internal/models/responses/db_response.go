package response

import (
	"net/http"

	apiErrors "github.com/TSI-Projects/group-project/internal/errors"
	"github.com/TSI-Projects/group-project/internal/repository"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
)

type DBResponse[T any] struct {
	dataArr []*T
	data    *T
	*BaseResponse

	dataArrName string
	dataName    string
}

func NewCustomerResponse[T any]() *DBResponse[T] {
	return &DBResponse[T]{BaseResponse: &BaseResponse{}}
}

func NewCustomerResponseWithError[T any](code, message string) ([]byte, error) {
	return NewCustomerResponse[T]().SetError(code, message).ToJSON()
}

func NewCustomerResponseWithData[T any](name string, data *T) ([]byte, error) {
	return NewCustomerResponse[T]().SetCustomer(data).SetDataName(name).SetSuccessStatus(true).ToJSON()
}

func NewCustomerResponseWithDataArr[T any](name string, dataArr []*T) ([]byte, error) {
	return NewCustomerResponse[T]().SetCustomers(dataArr).SetDataArrayName(name).SetSuccessStatus(true).ToJSON()
}

func (r *DBResponse[T]) SetError(code, message string) *DBResponse[T] {
	r.BaseResponse.NewError(code, message)
	return r
}

func (r *DBResponse[T]) SetCustomer(data *T) *DBResponse[T] {
	r.data = data
	return r
}

func (r *DBResponse[T]) SetCustomers(dataArr []*T) *DBResponse[T] {
	r.dataArr = dataArr
	return r
}

func (r *DBResponse[T]) SetSuccessStatus(isSuccess bool) *DBResponse[T] {
	r.BaseResponse.SetSuccessStatus(isSuccess)
	return r
}

func (r *DBResponse[T]) SetDataArrayName(name string) *DBResponse[T] {
	r.dataArrName = name
	return r
}

func (r *DBResponse[T]) SetDataName(name string) *DBResponse[T] {
	r.dataName = name
	return r
}

func (r *DBResponse[T]) ToJSON() ([]byte, error) {
	result := map[string]interface{}{
		"success": r.BaseResponse.Success,
	}

	if r.BaseResponse.Error != nil {
		result["error"] = r.BaseResponse.Error
	}

	if r.dataArrName != "" && r.dataArr != nil {
		result[r.dataArrName] = r.dataArr
	}

	if r.dataName != "" && r.data != nil {
		result[r.dataName] = r.data
	}

	return json.Marshal(result)
}

func WriteResponseWithError(w http.ResponseWriter, errorCode, errorMsg string, statusCode int) {
	resp, err := NewCustomerResponseWithError[repository.Customer](errorCode, errorMsg)
	if err != nil {
		log.Errorf("Failed to create error response: %v", err)
		http.Error(w, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}
	WriteResponse(w, statusCode, resp)
}

func WriteResponseWithSuccess(w http.ResponseWriter) {
	resp, err := NewCustomerResponse[string]().SetSuccessStatus(true).ToJSON()
	if err != nil {
		log.Errorf("failed to create success response: %v", err)
		http.Error(w, apiErrors.INTERNAL_ERROR_MESSAGE, http.StatusInternalServerError)
		return
	}
	WriteResponse(w, http.StatusOK, resp)
}

func WriteResponse(w http.ResponseWriter, statusCode int, resp []byte) {
	w.WriteHeader(statusCode)
	if _, err := w.Write(resp); err != nil {
		log.Errorf("Failed to write response: %v", err)
	}
}
