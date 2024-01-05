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

	resp, err := client.ConfirmPayment(&rozetkapay.ConfirmPaymentParams{
		ExternalID: "06e06b1e-0aeb-42fb-b9d9-40b6cc749d1f",
		Amount:     2000,
		Currency:   "UAH",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("IsSended: ", resp.IsSuccess)
}
