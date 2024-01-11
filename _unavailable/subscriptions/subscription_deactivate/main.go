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

	resp, err := client.DeactivateSubscription("subscriptionID")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message: ", resp.Message)
}
