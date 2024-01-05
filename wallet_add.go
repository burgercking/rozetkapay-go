package rozetkapay

import (
	"net/http"
	"time"
)

type AddWalletCustomerParams struct {
	CallbackURL   string        `json:"callback_url"`
	ResultURL     string        `json:"result_url"`
	PaymentMethod PaymentMethod `json:"payment_method" validate:"required"`
}

type (
	AddWalletCustomerPaymentMethod struct {
		Card     Card   `json:"card"`
		OptionID string `json:"option_id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	}
)
type AddWalletCustomerResponse struct {
	Action         PaymentResponseAction          `json:"action"`
	ActionRequired bool                           `json:"action_required"`
	CreatedAt      time.Time                      `json:"created_at"`
	PaymentMethod  AddWalletCustomerPaymentMethod `json:"payment_method"`
	Status         PaymentStatus                  `json:"status"`
}

// Adds new payment method to wallet. Either X-CUSTOMER-AUTH header or external_id param is required.
func (c *Client) AddWalletCustomerPayment(customerID string, params *AddWalletCustomerParams) (*AddWalletCustomerResponse, error) {
	req, err := c.NewRequest(http.MethodPost, BaseURL+"customers/v1/wallet", params, map[string]string{"external_id": customerID})
	if err != nil {
		return nil, err
	}

	resp := &AddWalletCustomerResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
