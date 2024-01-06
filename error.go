package rozetkapay

import (
	"errors"
	"fmt"
)

var (
	ErrFailedToValidate error = errors.New("failed to validate")
)

// ErrorResponse struct represents the error response structure.
type ErrorResponse struct {
	Code      PaymentStatusCode `json:"code"`
	Message   string            `json:"message"`
	Param     string            `json:"param"`
	PaymentID string            `json:"payment_id"`
	Type      string            `json:"type"`
}

// ErrorCode returns the error message.
func (e *ErrorResponse) ErrorCode() error {
	return fmt.Errorf(string(e.Code))
}
