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

	payments, err := client.GetSubscriptionPayments("subscriptionID")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscription payments amount: ", payments.Details.Amount)
}
