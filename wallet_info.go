package rozetkapay

import (
	"net/http"
	"time"
)

type (
	Card struct {
		ExpiresAt time.Time `json:"expires_at"`
		Mask      string    `json:"mask"`
	}

	WalletEntry struct {
		Card     Card   `json:"card"`
		OptionID string `json:"option_id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	}
)

type GetWalletInfoResponse struct {
	Address    string        `json:"address"`
	City       string        `json:"city"`
	Country    string        `json:"country"`
	Email      string        `json:"email"`
	ExternalID string        `json:"external_id"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	Patronym   string        `json:"patronym"`
	Phone      string        `json:"phone"`
	PostalCode string        `json:"postal_code"`
	RID        string        `json:"rid"`
	Wallet     []WalletEntry `json:"wallet"`
}

// Returns customer details including payment methods, if saved. Either X-CUSTOMER-AUTH header or external_id param is required.
func (c *Client) GetWalletCustomerPaymentInfo(customerID string) (*GetWalletInfoResponse, error) {
	req, err := c.NewRequest(http.MethodGet, BaseURL+"customers/v1/wallet", nil, map[string]string{"external_id": customerID})
	if err != nil {
		return nil, err
	}

	resp := &GetWalletInfoResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
