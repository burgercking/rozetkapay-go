package rozetkapay

import (
	"fmt"
	"net/http"
)

// https://cdn-epdev.rozetkapay.com/public-docs/index.html#tag/payments/operation/refundPayment
type RefundPaymentParams struct {
	ExternalID  string  `json:"external_id" validate:"required"`
	Amount      float64 `json:"amount,omitempty" validate:"required"`
	CallbackURL string  `json:"callback_url,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Payload     string  `json:"payload,omitempty"`
}

// Refund one-step payment after withdrawal, or two-step payment after confirmation.
func (c *Client) RefundPayment(params *RefundPaymentParams) (*PaymentResponse, error) {
	err := c.Validator.Struct(params)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodPost, BaseURL+"payments/v1/refund", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
