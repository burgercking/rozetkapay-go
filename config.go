package rozetkapay

import (
	"encoding/base64"
)

const (
	ProdStagingURL     = "https://api.rozetkapay.com/api/"
	DevStagingURL      = "https://api-epdev.rozetkapay.com/api/"
	DevStagingLogin    = "a6a29002-dc68-4918-bc5d-51a6094b14a8"
	DevStagingPassword = "XChz3J8qrr"
)

type Config struct {
	BasicAuth   string
	ResultURL   string
	CallbackURL string
	APIURL      string
}

func NewConfig(login, password string) *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString([]byte(login + ":" + password)),
		APIURL:    ProdStagingURL,
	}
}

func NewDevelopmentConfig() *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString([]byte(DevStagingLogin + ":" + DevStagingPassword)),
		// Dev is not working - code: authorization_failed
		APIURL: ProdStagingURL,
	}
}

func (c *Config) SetCallbackURL(callbackURL string) *Config {
	c.CallbackURL = callbackURL
	return c
}

func (c *Config) SetResultURL(resultURL string) *Config {
	c.ResultURL = resultURL
	return c
}

func (c *Config) SetAPI(url string) *Config {
	c.APIURL = url
	return c
}
