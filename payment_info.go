package rozetkapay

import (
	"net/http"
	"time"
)

type (
	Fee struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	}

	CancellationDetail struct {
		Amount            string            `json:"amount"`
		BillingOrderID    string            `json:"billing_order_id"`
		CreatedAt         time.Time         `json:"created_at"`
		Currency          string            `json:"currency"`
		Description       string            `json:"description"`
		GatewayOrderID    string            `json:"gateway_order_id"`
		Payload           string            `json:"payload"`
		PaymentID         string            `json:"payment_id"`
		ProcessedAt       time.Time         `json:"processed_at"`
		Properties        map[string]string `json:"properties"`
		RRN               string            `json:"rrn"`
		Status            string            `json:"status"`
		StatusCode        string            `json:"status_code"`
		StatusDescription string            `json:"status_description"`
		TransactionID     string            `json:"transaction_id"`
		AuthCode          string            `json:"auth_code"`
		Fee               Fee               `json:"fee"`
		TerminalName      string            `json:"terminal_name"`
	}

	ConfirmationDetail struct {
		Amount            string            `json:"amount"`
		BillingOrderID    string            `json:"billing_order_id"`
		CreatedAt         time.Time         `json:"created_at"`
		Currency          string            `json:"currency"`
		Description       string            `json:"description"`
		GatewayOrderID    string            `json:"gateway_order_id"`
		Payload           string            `json:"payload"`
		PaymentID         string            `json:"payment_id"`
		ProcessedAt       time.Time         `json:"processed_at"`
		Properties        map[string]string `json:"properties"`
		RRN               string            `json:"rrn"`
		Status            string            `json:"status"`
		StatusCode        string            `json:"status_code"`
		StatusDescription string            `json:"status_description"`
		TransactionID     string            `json:"transaction_id"`
		AuthCode          string            `json:"auth_code"`
		Fee               Fee               `json:"fee"`
		TerminalName      string            `json:"terminal_name"`
	}

	PurchaseDetail struct {
		Amount            string            `json:"amount"`
		BillingOrderID    string            `json:"billing_order_id"`
		CreatedAt         time.Time         `json:"created_at"`
		Currency          string            `json:"currency"`
		Description       string            `json:"description"`
		GatewayOrderID    string            `json:"gateway_order_id"`
		Payload           string            `json:"payload"`
		PaymentID         string            `json:"payment_id"`
		ProcessedAt       time.Time         `json:"processed_at"`
		Properties        map[string]string `json:"properties"`
		RRN               string            `json:"rrn"`
		Status            string            `json:"status"`
		StatusCode        string            `json:"status_code"`
		StatusDescription string            `json:"status_description"`
		TransactionID     string            `json:"transaction_id"`
		AuthCode          string            `json:"auth_code"`
		Fee               Fee               `json:"fee"`
		TerminalName      string            `json:"terminal_name"`
	}

	RefundDetail struct {
		Amount            string            `json:"amount"`
		BillingOrderID    string            `json:"billing_order_id"`
		CreatedAt         time.Time         `json:"created_at"`
		Currency          string            `json:"currency"`
		Description       string            `json:"description"`
		GatewayOrderID    string            `json:"gateway_order_id"`
		Payload           string            `json:"payload"`
		PaymentID         string            `json:"payment_id"`
		ProcessedAt       time.Time         `json:"processed_at"`
		Properties        map[string]string `json:"properties"`
		RRN               string            `json:"rrn"`
		Status            string            `json:"status"`
		StatusCode        string            `json:"status_code"`
		StatusDescription string            `json:"status_description"`
		TransactionID     string            `json:"transaction_id"`
		AuthCode          string            `json:"auth_code"`
		Fee               Fee               `json:"fee"`
		TerminalName      string            `json:"terminal_name"`
	}
)

// https://cdn-epdev.rozetkapay.com/public-docs/index.html#tag/payments/operation/paymentInfo
type PaymentInfoResponse struct {
	Action              PaymentResponseAction `json:"action"`
	ActionRequired      bool                  `json:"action_required"`
	Amount              string                `json:"amount"`
	AmountCanceled      string                `json:"amount_canceled"`
	AmountConfirmed     string                `json:"amount_confirmed"`
	AmountRefunded      string                `json:"amount_refunded"`
	Canceled            bool                  `json:"canceled"`
	CancellationDetails []CancellationDetail  `json:"cancellation_details"`
	ConfirmationDetails []ConfirmationDetail  `json:"confirmation_details"`
	Confirmed           bool                  `json:"confirmed"`
	CreatedAt           time.Time             `json:"created_at"`
	Currency            string                `json:"currency"`
	ExternalID          string                `json:"external_id"`
	ID                  string                `json:"id"`
	PurchaseDetails     []PurchaseDetail      `json:"purchase_details"`
	Purchased           bool                  `json:"purchased"`
	ReceiptURL          string                `json:"receipt_url"`
	RefundDetails       []RefundDetail        `json:"refund_details"`
	Refunded            bool                  `json:"refunded"`
	Customer            Recipient             `json:"customer"`
}

// Get payment info by id.
func (c *Client) GetPaymentInfo(externalID string) (*PaymentInfoResponse, error) {
	req, err := c.NewRequest(http.MethodGet, BaseURL+"payments/v1/info", nil, map[string]string{"external_id": externalID})
	if err != nil {
		return nil, err
	}

	resp := &PaymentInfoResponse{}
	if err := c.Send(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
