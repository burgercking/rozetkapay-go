package main

import (
	"log"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
)

func main() {
	var (
		cfg    = rozetkapay.NewDevelopmentConfig().SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	resp, err := client.AddWalletCustomerPayment("123123", &rozetkapay.AddWalletCustomerSchema{
		CallbackURL: cfg.CallbackURL,
		PaymentMethod: rozetkapay.PaymentMethod{
			Type: rozetkapay.PaymentMethodTypeWallet,
			Wallet: rozetkapay.Wallet{
				BrowserFingerprint: rozetkapay.BrowserFingerprint{},
				OptionID:           "1332",
				Use3DSFlow:         true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Type: ", resp.Action.Type)
	log.Println("Value: ", resp.Action.Value)
}
