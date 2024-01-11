package main

import (
	"log"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
)

func main() {
	var (
		cfg     = rozetkapay.NewDevelopmentConfig().SetCallbackURL(examples.DevEnvironmentCallbackURL)
		manager = rozetkapay.NewManager(cfg)
	)

	resp, err := manager.RefundPayment(&rozetkapay.RefundPaymentSchema{
		ExternalID:  "5931a91a-33c3-4a3e-8856-6624b0c29fe2",
		Amount:      200,
		CallbackURL: cfg.CallbackURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status: ", resp.IsSuccess)
}
