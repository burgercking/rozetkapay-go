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

	plans, err := client.GetSubscriptionPlans("prom")
	if err != nil {
		log.Fatal(err)
	}

	for _, plan := range plans {
		log.Println("Plan name: ", plan.Name)
	}
}
