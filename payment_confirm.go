package rozetkapay

import (
	"fmt"
	"net/http"
)

// https://cdn.rozetkapay.com/public-docs/index.html#tag/payments/operation/confirmPayment
type ConfirmPaymentParams struct {
	ExternalID  string  `json:"external_id" validate:"required"`
	Amount      float64 `json:"amount,omitempty"`
	CallbackURL string  `json:"callback_url,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Payload     string  `json:"payload,omitempty"`
}

// Confirm two-step payment.
func (c *Client) ConfirmPayment(params *ConfirmPaymentParams) (*PaymentResponse, error) {
	err := c.Validator.Struct(params)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodPost, BaseURL+"payments/v1/confirm", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
