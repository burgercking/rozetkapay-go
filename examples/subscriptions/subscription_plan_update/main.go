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

	resp, err := client.UpdateSubscriptionPlan("planID", &rozetkapay.UpdateSubscriptionPlanSchema{
		Name: "VIP+",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message: ", resp.Message)
}
