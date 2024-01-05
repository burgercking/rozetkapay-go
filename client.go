package rozetkapay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const (
	// BaseURL string = "https://api-epdev.rozetkapay.com/api/"
	BaseURL string = "https://api.rozetkapay.com/api/"
	// GatewayVersion string = "0.1"
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
		if err := json.Unmarshal(body, &errResp); err != nil {
			return err
		}
		return errResp.Error()
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
