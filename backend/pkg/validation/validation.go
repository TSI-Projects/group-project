package validation

import (
	"fmt"
	"strings"

	"github.com/TSI-Projects/group-project/utils"
	"github.com/go-playground/validator/v10"
)

type ValidatorClient struct {
	validator *validator.Validate
}

func NewValidatorClient() *ValidatorClient {
	return &ValidatorClient{
		validator: validator.New(),
	}
}

func (c *ValidatorClient) Validate(data any) error {
	if data == nil {
		return fmt.Errorf("validation failed: data is nil")
	}

	if err := c.validator.Struct(data); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			missingFields := make([]string, 0, len(validationErrors))

			for _, validationErr := range validationErrors {
				field := utils.SplitOnUppercase(validationErr.Field())
				missingFields = append(missingFields, field)
			}

			return fmt.Errorf("required field is not received: %s", strings.Join(missingFields, ", "))
		}
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
