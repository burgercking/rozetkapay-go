package rozetkapay

import (
	"fmt"
	"net/http"
)

// https://cdn-epdev.rozetkapay.com/public-docs/index.html#tag/payments/operation/createPayment

type CustomerColorModer string

const (
	CustomerColorModeLight CustomerColorModer = "light"
	CustomerColorModeDark  CustomerColorModer = "dark"
)

type CustomerCheckoutLocale string

const (
	CustomerCheckoutLocaleUK CustomerCheckoutLocale = "UK"
	CustomerCheckoutLocaleEN CustomerCheckoutLocale = "EN"
	CustomerCheckoutLocaleES CustomerCheckoutLocale = "ES"
	CustomerCheckoutLocalePL CustomerCheckoutLocale = "PL"
	CustomerCheckoutLocaleFR CustomerCheckoutLocale = "FR"
	CustomerCheckoutLocaleSK CustomerCheckoutLocale = "SK"
	CustomerCheckoutLocaleDE CustomerCheckoutLocale = "DE"
)

type PaymentMethodType string

const (
	PaymentMethodTypeApplePay  PaymentMethodType = "apple_pay"
	PaymentMethodTypeCCToken   PaymentMethodType = "cc_token"
	PaymentMethodTypeGooglePay PaymentMethodType = "google_pay"
	PaymentMethodTypeWallet    PaymentMethodType = "wallet"
)

type (
	CustomerData struct {
		// Available values: CustomerColorModeLight - "light"; CustomerColorModeDark = "dark".
		// Checkout theme for hosted type of interaction.
		// If the parameter is not filled in, the default theme will be selected.
		ColorMode CustomerColorModer `json:"color_mode,omitempty"`

		// Available values:
		// CustomerCheckoutLocaleUK` = "UK"; CustomerCheckoutLocaleEN = "EN";
		// CustomerCheckoutLocaleES = "ES"; CustomerCheckoutLocalePL = "PL";
		// CustomerCheckoutLocaleFR = "FR"; CustomerCheckoutLocaleSK = "SK";
		// CustomerCheckoutLocaleDE = "DE".
		// Checkout locale for hosted type of interaction.
		// If the field is not filled in, the browser locale or the default locale will be selected.
		Locale CustomerCheckoutLocale `json:"locale,omitempty"`

		// Personal account number when refilling services.
		AccountNumber string `json:"account_number,omitempty"`
		IPAddress     string `json:"ip_address,omitempty"`
		Address       string `json:"address,omitempty" validate:"lte=50"`
		City          string `json:"city,omitempty"`
		Country       string `json:"country,omitempty"`
		Email         string `json:"email,omitempty"`

		// Unique payer number of the partner.
		ExternalID string `json:"external_id,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Patronym   string `json:"patronym,omitempty"`

		// Block for selecting the payer's payment method.
		// The field is required for the direct integration method.
		PaymentMethod PaymentMethod `json:"payment_method,omitempty" validate:"required"`
		Phone         string        `json:"phone,omitempty"`
		PostalCode    string        `json:"postal_code,omitempty"`
	}
)

type CreatePaymentParams struct {
	// Amount of the order.
	Amount float64 `json:"amount" validate:"required"`

	// Currency of the order (ISO 4207).
	Currency string `json:"currency" validate:"required"`

	// Unique order number.
	ExternalID string `json:"external_id" validate:"required"`

	// Type of interaction, possible values:
	// "hosted" - returns a link to the payment page;
	// "direct" - direct host2host interaction.
	//  When mode is set to direct - customer field becomes required.
	Mode PaymentMode `json:"mode" validate:"required"`

	// Address where the callback with the status of the operation will be sent.
	// Pass this only if you have not already set it in confg.
	CallbackURL string `json:"callback_url,omitempty"`

	// The user will be redirected to this address after payment.
	// Pass this only if you have not already set it in confg.
	ResultURL string `json:"result_url,omitempty"`

	// "True" - the funds will be debited immediately after the payment is made;
	// "False" - to make a write-off, you need to additionally call the confirm payment method.
	// https://docs.google.com/document/d/1AbNXlJlPdzjZcpotd83Qb7GWXt78UhYGRY-GQRWI35M/edit#heading=h.n5tet6vz5p3g
	Confirm bool `json:"confirm,omitempty"`

	// Description of the order.
	Description string `json:"description,omitempty"`

	// Field for additional data. Maximum length 4000 characters
	Payload string `json:"payload,omitempty" validate:"max=4000"`

	// Payer information block.
	Customer *CustomerData `json:"customer" validate:"required_if=Mode direct"`

	// List of products/services in the order.
	Products []Product `json:"products,omitempty"`

	// A block of data about the recipient of funds. Used for p2p transactions.
	Recipient *Recipient `json:"recipient,omitempty"`

	// Dictionary in key:value format for additional parameters.
	Properties map[string]string `json:"properties,omitempty"`
}

// Creates payment and performs desired operation.
func (c *Client) CreatePayment(params *CreatePaymentParams) (*PaymentResponse, error) {
	err := c.Validator.Struct(params)
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	req, err := c.NewRequest(http.MethodPost, BaseURL+"payments/v1/new", params, nil)
	if err != nil {
		return nil, err
	}

	resp := &PaymentResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
