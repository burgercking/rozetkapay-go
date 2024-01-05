package main

import (
	"log"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
)

func main() {
	var (
		cfg = rozetkapay.NewConfig(rozetkapay.DevEnvironmentProjectID, rozetkapay.DevEnvironmentPassword).
			SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	canceled, err := client.CancelPayment(&rozetkapay.CancelPaymentParams{
		ExternalID:  "b9729678-a0fa-47d6-8171-2ea638f31dfa",
		CallbackURL: cfg.CallbackURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status: ", canceled)
}
