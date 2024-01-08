package main

import (
	"log"
	"time"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
)

func main() {
	var (
		cfg    = rozetkapay.NewDevelopmentConfig().SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	plan, err := client.CreateSubscriptionPlan(&rozetkapay.CreateSubscriptionPlanSchema{
		Name:          "VIP",
		Description:   "Get access to additional features with the VIP subscription!",
		Price:         200,
		Currency:      "UAH",
		Platforms:     []string{},
		FrequencyType: rozetkapay.FrequencyTypeMonthly,
		Frequency:     1,
		StartDate:     time.Now().String(),
		EndDate:       time.Now().Add(24 * time.Hour).String(),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Plan ID: ", plan.ID)
	log.Println("Plan State: ", plan.State)
	log.Println("Plan CreatedAt: ", plan.CreatedAt)
}
