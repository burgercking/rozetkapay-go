package rozetkapay

import (
	"encoding/base64"
)

const (
	DevEnvironmentProjectID = "a6a29002-dc68-4918-bc5d-51a6094b14a8"
	DevEnvironmentPassword  = "XChz3J8qrr"
)

type Config struct {
	BasicAuth   string
	ResultURL   string
	CallbackURL string
}

func NewConfig(projectID, password string) *Config {
	return &Config{
		BasicAuth: base64.StdEncoding.EncodeToString([]byte(projectID + ":" + password)),
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
