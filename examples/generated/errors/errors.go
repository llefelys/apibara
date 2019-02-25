package errors

import "fmt"

type (
	minValidationError struct {
		param string
		min   int
	}
	maxValidationError struct {
		param string
		max   int
	}
	requiredValidationError struct {
		param string
	}
	posibleValuesValidationError struct {
		param string
	}
	badRequestError struct {
		param string
		err   error
	}
)

func (m minValidationError) Error() string {
	return fmt.Sprintf("parameter %s can't be lower then %d", m.param, m.min)
}
func (m maxValidationError) Error() string {
	return fmt.Sprintf("parameter %s can't be bigger then %d", m.param, m.max)
}
func (m requiredValidationError) Error() string {
	return fmt.Sprintf("parameter %s is required", m.param)
}
func (m posibleValuesValidationError) Error() string {
	return fmt.Sprintf("parameter %s is not allowed", m.param)
}
func (m badRequestError) Error() string {
	return fmt.Sprintf("error durring parameter %s parsing: %s", m.param, m.err)
}
func NewMinValidationError(param string, min int) error {
	return minValidationError{param: param, min: min}
}
func NewMaxValidationError(param string, min int) error {
	return maxValidationError{param: param, max: min}
}
func NewRequiredValidationError(param string) error { return requiredValidationError{param: param} }
func NewPosibleValuesValidationError(param string) error {
	return posibleValuesValidationError{param: param}
}
func NewBadRequestError(param string, err error) error { return badRequestError{param: param, err: err} }
