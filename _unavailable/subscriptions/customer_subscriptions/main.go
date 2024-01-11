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

	subs, err := client.GetCustomerSubscriptions("customerID")
	if err != nil {
		log.Fatal(err)
	}

	for _, sub := range subs {
		log.Println("Subscription ID: ", sub.ID)
		log.Println("Subscription State: ", sub.State)
		log.Println("Subscription CreatedAt: ", sub.CreatedAt)
	}
}
