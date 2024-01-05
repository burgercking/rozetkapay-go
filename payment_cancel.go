package rozetkapay

import (
	"fmt"
	"net/http"
)

// https://cdn-epdev.rozetkapay.com/public-docs/index.html#tag/payments/operation/cancelPayment
type CancelPaymentParams struct {
	ExternalID  string  `json:"external_id" validate:"required"`
	Amount      float64 `json:"amount,omitempty"`
	CallbackURL string  `json:"callback_url,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Payload     string  `json:"payload,omitempty"`
}

// Cancel two-step payment.
func (c *Client) CancelPayment(params *CancelPaymentParams) (*PaymentResponse, error) {
	err := c.Validator.Struct(params)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodPost, BaseURL+"payments/v1/cancel", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
