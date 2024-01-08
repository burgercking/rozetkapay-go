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

	plan, err := client.GetSubscriptionPlan("planID")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Plan ID: ", plan.ID)
	log.Println("Plan State: ", plan.State)
	log.Println("Plan CreatedAt: ", plan.CreatedAt)
}
