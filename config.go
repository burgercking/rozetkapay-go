package rozetkapay

import (
	"encoding/base64"
)

const (
	API_URL     = "https://api.rozetkapay.com/api/"
	DevLogin    = "a6a29002-dc68-4918-bc5d-51a6094b14a8"
	DevPassword = "XChz3J8qrr"
)

type Config struct {
	API         string
	BasicAuth   string
	ResultURL   string
	CallbackURL string
	IsDebug     bool
}

func NewConfig(login, password string) *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString(
			[]byte(login + ":" + password),
		),
		API: API_URL,
	}
}

func NewDevelopmentConfig() *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString(
			[]byte(DevLogin + ":" + DevPassword),
		),
		API:     API_URL,
		IsDebug: true,
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

func (c *Config) SetDebugMode(debug bool) *Config {
	c.IsDebug = debug
	return c
}
