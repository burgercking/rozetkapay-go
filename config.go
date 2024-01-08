package rozetkapay

import (
	"encoding/base64"
)

const (
	APIURL      = "https://api.rozetkapay.com/api/"
	DevLogin    = "a6a29002-dc68-4918-bc5d-51a6094b14a8"
	DevPassword = "XChz3J8qrr"
)

type Config struct {
	BasicAuth   string
	ResultURL   string
	CallbackURL string
	APIURL      string
}

func NewConfig(login, password string) *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString(
			[]byte(login + ":" + password),
		),
		APIURL: APIURL,
	}
}

func NewDevelopmentConfig() *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString(
			[]byte(DevLogin + ":" + DevPassword),
		),
		APIURL: APIURL,
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
