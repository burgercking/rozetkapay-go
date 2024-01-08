package rozetkapay

import (
	"fmt"
	"net/http"
)

// Adds new payment method to wallet.
func (c *Client) AddWalletCustomerPayment(customerID string, schema *AddWalletCustomerSchema) (
	*AddWalletCustomerResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"customers/v1/wallet",
		schema,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &AddWalletCustomerResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Returns customer details including payment methods, if saved.
func (c *Client) GetWalletCustomerPaymentInfo(customerID string) (
	*GetWalletInfoResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"customers/v1/wallet",
		nil,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &GetWalletInfoResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Deletes customer payment method from wallet.
func (c *Client) DeleteWalletCustomerPayment(customerID string, schema *DeleteWalletCustomerSchema) (
	*DeleteWalletCustomerResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodDelete,
		c.Config.APIURL+"customers/v1/wallet",
		schema,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &DeleteWalletCustomerResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
