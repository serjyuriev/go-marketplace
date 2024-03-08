package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductValidationError(t *testing.T) {
	tests := []struct {
		name  string
		field string
		err   error
	}{
		{
			name:  "passing non nil error",
			field: "SomeField",
			err:   errors.New("sample error text"),
		},
		{
			name:  "passing nil error",
			field: "SomeField",
			err:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationError := NewProductValidationError(tt.field, tt.err)

			assert.Equal(t, tt.field, validationError.Field)
			assert.Equal(t, tt.err, validationError.Err)
		})
	}
}

func TestProductValidationError_Error(t *testing.T) {
	tests := []struct {
		name        string
		err         *ProductValidationError
		expectedMsg string
	}{
		{
			name: "with non nil inner error",
			err: &ProductValidationError{
				Field: "SomeField",
				Err:   errors.New("sample error text"),
			},
			expectedMsg: "Field: SomeField, did not passed validation check (sample error text)",
		},
		{
			name: "with nil inner error",
			err: &ProductValidationError{
				Field: "SomeField",
				Err:   nil,
			},
			expectedMsg: "Field: SomeField, did not passed validation check",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := tt.err.Error()

			assert.Equal(t, tt.expectedMsg, msg)
		})
	}
}
