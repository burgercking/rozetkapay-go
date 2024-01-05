package rozetkapay

import (
	"fmt"
	"net/http"
)

type CallbackResendOperation string

const (
	CallbackResendOperationPayment CallbackResendOperation = "payment"
)

type PaymentCallbackResendParams struct {
	ExternalID string `json:"external_id" validate:"required"`

	// Payment operation possible values:
	// CallbackResendOperationPayment - "payment".
	Operation CallbackResendOperation `json:"operation"`
}

// Prepares the data about the specified payment of transaction and sends it into callback_url which was provided on the payment step.
// If the operation field is not provided the callback will be sent for the last operation.
func (c *Client) ResendPaymentCallback(params *PaymentCallbackResendParams) (resended bool, err error) {
	err = c.Validator.Struct(params)
	if err != nil {
		return false, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodPost, BaseURL+"payments/v1/callback/resend", params, nil)
	if err != nil {
		return false, err
	}

	if err := c.Send(req, nil); err != nil {
		return false, err
	}

	return true, nil
}
