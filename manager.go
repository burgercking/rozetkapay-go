package rozetkapay

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

type Manager struct {
	Config *Config
	Client *http.Client
}

func NewManager(config *Config, opts ...ManagerOpts) *Manager {
	m := &Manager{
		Config: config,
		Client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

type ManagerOpts func(*Manager)

func WithCustomClient(c *http.Client) ManagerOpts {
	return func(m *Manager) {
		m.Client = c
	}
}

func (m *Manager) Send(req *http.Request, v interface{}) error {
	req.Header = http.Header{
		"Content-type":  {"application/json"},
		"Authorization": {"Basic " + m.Config.BasicAuth},
	}

	if m.Config.IsDebug {
		slog.Info("RozetkaPay",
			slog.String("type", "request"),
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
		)
	}

	resp, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errResp *ErrorResponse
		if len(body) == 0 {
			return ErrResponseIsEmpty
		}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return err
		}

		if m.Config.IsDebug {
			slog.Error("RozetkaPay error",
				slog.String("type", "response"),
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
				slog.Any("code", errResp.Code),
				slog.String("message", errResp.Message),
				slog.String("payment_id", errResp.PaymentID),
				slog.String("type", errResp.Type),
			)
		}

		return errResp.ErrorCode()
	}

	if v == nil {
		return nil
	}

	if m.Config.IsDebug {
		slog.Info("RozetkaPay",
			slog.String("type", "response"),
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
			slog.Int("code", resp.StatusCode),
			slog.Int("bytes", len(body)),
		)
	}

	err = json.Unmarshal(body, v)
	return err
}

func (m *Manager) NewRequest(method, url string, payload interface{}, query map[string]string) (
	*http.Request, error,
) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Creates payment and performs desired operation.
func (m *Manager) CreatePayment(schema *CreatePaymentSchema) (
	*PaymentResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"payments/v1/new",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Confirm two-step payment.
func (m *Manager) ConfirmPayment(schema *ConfirmPaymentSchema) (
	*PaymentResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"payments/v1/confirm",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Cancel two-step payment.
func (m *Manager) CancelPayment(schema *CancelPaymentSchema) (
	*PaymentResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"payments/v1/cancel",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Refund one-step payment after withdrawal, or two-step payment after confirmation.
func (m *Manager) RefundPayment(schema *RefundPaymentSchema) (
	*PaymentResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"payments/v1/refund",
		schema,
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Get payment info by id.
func (m *Manager) GetPaymentInfo(externalID string) (
	*PaymentInfoResponse, error,
) {
	req, err := m.NewRequest(
		http.MethodGet,
		m.Config.API+"payments/v1/info",
		nil,
		map[string]string{"external_id": externalID},
	)
	if err != nil {
		return nil, err
	}
	resp := &PaymentInfoResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Parsing callback from the body.
func (m *Manager) GetPaymentCallbackFromBytes(body []byte) (
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
func (m *Manager) ResendPaymentCallback(schema *PaymentCallbackResendSchema) (
	resended bool, err error,
) {
	err = v.Struct(schema)
	if err != nil {
		return false, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"payments/v1/callback/resend",
		schema,
		nil,
	)
	if err != nil {
		return false, err
	}
	if err := m.Send(req, nil); err != nil {
		return false, err
	}
	return true, nil
}

// Adds new payment method to wallet.
func (m *Manager) AddWalletCustomerPayment(customerID string, schema *AddWalletCustomerSchema) (
	*AddWalletCustomerResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodPost,
		m.Config.API+"customers/v1/wallet",
		schema,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &AddWalletCustomerResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Returns customer details including payment methods, if saved.
func (m *Manager) GetWalletCustomerPaymentInfo(customerID string) (
	*GetWalletInfoResponse, error,
) {
	req, err := m.NewRequest(
		http.MethodGet,
		m.Config.API+"customers/v1/wallet",
		nil,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &GetWalletInfoResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Deletes customer payment method from wallet.
func (m *Manager) DeleteWalletCustomerPayment(customerID string, schema *DeleteWalletCustomerSchema) (
	*DeleteWalletCustomerResponse, error,
) {
	err := v.Struct(schema)
	if err != nil {
		return nil, err
	}
	req, err := m.NewRequest(
		http.MethodDelete,
		m.Config.API+"customers/v1/wallet",
		schema,
		map[string]string{"external_id": customerID},
	)
	if err != nil {
		return nil, err
	}
	resp := &DeleteWalletCustomerResponse{}
	if err := m.Send(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
