package validator_test

import (
	"testing"

	"github.com/TSI-Projects/group-project/pkg/validation"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,gte=18"`
}

func TestValidatorClient_Validate_Valid(t *testing.T) {
	validatorClient := validation.NewValidatorClient()
	data := TestStruct{
		Name:  "John Doe",
		Email: "test@example.com",
		Age:   20,
	}

	err := validatorClient.Validate(data)
	assert.NoError(t, err, "Validation should pass without any errors")
}

func TestValidatorClient_Validate_MissingFields(t *testing.T) {
	validatorClient := validation.NewValidatorClient()
	data := TestStruct{}

	err := validatorClient.Validate(data)
	assert.Error(t, err, "expected error for empty fields")

	if err != nil {
		expectedMsg := "required field is not received"
		assert.Contains(t, err.Error(), expectedMsg, "expected text about required fields")

		assert.Contains(t, err.Error(), "Name", "expected 'Name' field to be in error")
		assert.Contains(t, err.Error(), "Email", "expected 'Email' field to be in error")
		assert.Contains(t, err.Error(), "Age", "expected 'Age' field to be in error")
	}
}

func TestValidatorClient_Validate_NilData(t *testing.T) {
	validatorClient := validation.NewValidatorClient()

	err := validatorClient.Validate(nil)
	assert.Error(t, err, "expected error for nil data")
}
