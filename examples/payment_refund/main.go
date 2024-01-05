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

	resp, err := client.RefundPayment(&rozetkapay.RefundPaymentParams{
		ExternalID:  "5931a91a-33c3-4a3e-8856-6624b0c29fe2",
		Amount:      200,
		CallbackURL: cfg.CallbackURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status: ", resp.IsSuccess)
}
