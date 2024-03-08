package errors

import "fmt"

// ProductValidationError представляет собой ошибку валидации одного из полей продукта.
type ProductValidationError struct {
	Field string
	Err   error
}

// NewProductValidationError создает новую ошибку с указанием поля, не прошедшего валидацию.
func NewProductValidationError(field string, err error) *ProductValidationError {
	return &ProductValidationError{Field: field, Err: err}
}

// Error реализует интерфейс error для ProductValidationError.
func (e *ProductValidationError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Field: %s, did not passed validation check", e.Field)
	} else {
		return fmt.Sprintf("Field: %s, did not passed validation check (%s)", e.Field, e.Err.Error())
	}
}
