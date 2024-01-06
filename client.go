package rozetkapay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Client struct {
	Config    *Config
	Client    *http.Client
	Validator *validator.Validate
}

func NewClient(config *Config) *Client {
	return &Client{
		Config:    config,
		Client:    &http.Client{},
		Validator: validator.New(),
	}
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	req.Header = http.Header{
		"Content-type":  {"application/json"},
		"Authorization": {"Basic " + c.Config.BasicAuth},
	}

	resp, err := c.Client.Do(req)
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
			return errors.New("[ROZETKAPAY] Error: empty, HTTP status code: " + resp.Status)
		}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return err
		}
		log.Printf(
			"[ROZETKAPAY] Error: [code=%s] [message=%s] [payment_id=%s] [type=%s]",
			errResp.Code, errResp.Message, errResp.PaymentID, errResp.Type,
		)
		return errResp.ErrorCode()

	} else if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) NewRequest(method, url string, payload interface{}, query map[string]string) (*http.Request, error) {
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
