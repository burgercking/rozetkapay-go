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

	sub, err := client.GetSubscriptionInfo("subscriptionID")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscription ID: ", sub.ID)
	log.Println("Subscription State: ", sub.State)
	log.Println("Subscription CreatedAt: ", sub.CreatedAt)
}
