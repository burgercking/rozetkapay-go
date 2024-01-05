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

	info, err := client.GetPaymentInfo("06e06b1e-0aeb-42fb-b9d9-40b6cc749d1f")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Amount: ", info.Amount)
	log.Println("Refunded: ", info.AmountRefunded)
	log.Println("Confirmed: ", info.Confirmed)
	log.Println("Canceled: ", info.Canceled)
}
