package rozetkapay

import "time"

// Add wallet customer
type (
	AddWalletCustomerSchema struct {
		CallbackURL   string        `json:"callback_url"`
		ResultURL     string        `json:"result_url"`
		PaymentMethod PaymentMethod `json:"payment_method" validate:"required"`
	}

	AddWalletCustomerPaymentMethod struct {
		Card     Card   `json:"card"`
		OptionID string `json:"option_id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	}

	AddWalletCustomerResponse struct {
		Action         PaymentUserAction              `json:"action"`
		ActionRequired bool                           `json:"action_required"`
		CreatedAt      time.Time                      `json:"created_at"`
		PaymentMethod  AddWalletCustomerPaymentMethod `json:"payment_method"`
		Status         PaymentStatus                  `json:"status"`
	}
)

// Get wallet info
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

	GetWalletInfoResponse struct {
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
)

// Delete wallet customer
type (
	DeleteWalletCustomerSchema struct {
		OptionID string            `json:"option_id" validate:"required"`
		Type     PaymentMethodType `json:"type" validate:"required"`
	}

	DeleteWalletCustomerResponse struct {
		Delete   bool              `json:"delete"`
		OptionID string            `json:"option_id"`
		Type     PaymentMethodType `json:"type"`
	}
)
