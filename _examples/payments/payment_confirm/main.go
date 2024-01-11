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

	resp, err := manager.ConfirmPayment(&rozetkapay.ConfirmPaymentSchema{
		ExternalID: "06e06b1e-0aeb-42fb-b9d9-40b6cc749d1f",
		Amount:     2000,
		Currency:   "UAH",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("IsSended: ", resp.IsSuccess)
}
