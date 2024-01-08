package rozetkapay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Creates payment and performs desired operation.
func (c *Client) CreatePayment(schema *CreatePaymentSchema) (
	*PaymentResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"payments/v1/new",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Confirm two-step payment.
func (c *Client) ConfirmPayment(schema *ConfirmPaymentSchema) (
	*PaymentResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"payments/v1/confirm",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Cancel two-step payment.
func (c *Client) CancelPayment(schema *CancelPaymentSchema) (
	*PaymentResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"payments/v1/cancel",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Refund one-step payment after withdrawal, or two-step payment after confirmation.
func (c *Client) RefundPayment(schema *RefundPaymentSchema) (
	*PaymentResponse, error,
) {
	err := c.Validator.Struct(schema)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"payments/v1/refund",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get payment info by id.
func (c *Client) GetPaymentInfo(externalID string) (
	*PaymentInfoResponse, error,
) {
	req, err := c.NewRequest(
		http.MethodGet,
		c.Config.APIURL+"payments/v1/info",
		nil,
		map[string]string{"external_id": externalID},
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentInfoResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Parsing callback from the body.
func (c *Client) GetPaymentCallbackFromBytes(body []byte) (
	*PaymentResponse, error,
) {
	var callback *PaymentResponse
	if err := json.Unmarshal(body, &callback); err != nil {
		return nil, err
	}
	return callback, nil
}

// Prepares the data about the specified payment of transaction and sends it into callback_url which was provided on the payment step.
// If the operation field is not provided the callback will be sent for the last operation.
func (c *Client) ResendPaymentCallback(schema *PaymentCallbackResendSchema) (
	resended bool, err error,
) {
	err = c.Validator.Struct(schema)
	if err != nil {
		return false, fmt.Errorf("%s: %w", ErrFailedToValidate, err)
	}
	req, err := c.NewRequest(
		http.MethodPost,
		c.Config.APIURL+"payments/v1/callback/resend",
		schema,
		nil,
	)
	if err != nil {
		return false, err
	}
	if err := c.Send(req, nil); err != nil {
		return false, err
	}
	return true, nil
}
