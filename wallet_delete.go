package rozetkapay

import (
	"fmt"
	"net/http"
)

type DeleteWalletCustomerParams struct {
	OptionID string            `json:"option_id" validate:"required"`
	Type     PaymentMethodType `json:"type" validate:"required"`
}

type DeleteWalletCustomerResponse struct {
	Delete   bool              `json:"delete"`
	OptionID string            `json:"option_id"`
	Type     PaymentMethodType `json:"type"`
}

// Deletes customer payment method from wallet. Either X-CUSTOMER-AUTH header or external_id param is required.
func (c *Client) DeleteWalletCustomerPayment(customerID string, params *DeleteWalletCustomerParams) (*DeleteWalletCustomerResponse, error) {
	err := c.Validator.Struct(params)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodDelete, BaseURL+"customers/v1/wallet", params, map[string]string{"external_id": customerID})
	if err != nil {
		return nil, err
	}

	resp := &DeleteWalletCustomerResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
