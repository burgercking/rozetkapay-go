package rozetkapay

import (
	"encoding/json"
)

// Parsing callback from the body.
func (c *Client) GetPaymentCallbackFromBytes(body []byte) (*PaymentResponse, error) {
	var callback *PaymentResponse
	if err := json.Unmarshal(body, &callback); err != nil {
		return nil, err
	}
	return callback, nil
}
