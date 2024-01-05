package rozetkapay

import "fmt"

// ErrorResponse struct represents the error response structure.
type ErrorResponse struct {
	Code      PaymentStatusCode `json:"code"`
	Message   string            `json:"message"`
	Param     string            `json:"param"`
	PaymentID string            `json:"payment_id"`
	Type      string            `json:"type"`
}

// Error returns the error message.
func (e *ErrorResponse) Error() error {
	return fmt.Errorf(
		"[ROZETKAPAY] Error: [code=%s] [message=%s] [payment_id=%s] [type=%s]",
		e.Code, e.Message, e.PaymentID, e.Type,
	)
}
