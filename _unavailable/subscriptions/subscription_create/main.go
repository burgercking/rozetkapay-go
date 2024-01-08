package main

import (
	"log"
	"time"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
	"github.com/google/uuid"
)

func main() {
	var (
		cfg    = rozetkapay.NewDevelopmentConfig().SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	resp, err := client.CreateSubscription(&rozetkapay.CreateSubscriptionSchema{
		AutoRenew:   true,
		CallbackURL: cfg.CallbackURL,
		Customer: rozetkapay.Recipient{
			Address:    "sss",
			City:       "sssss",
			Country:    "sss",
			Email:      "seeee@gmail.com",
			ExternalID: uuid.New().String(),
			FirstName:  "John",
			LastName:   "Doe",
			Patronym:   "Dovich",
			PaymentMethod: rozetkapay.PaymentMethod{
				Type: rozetkapay.PaymentMethodTypeWallet,
				Wallet: rozetkapay.Wallet{
					OptionID:   uuid.New().String(),
					Use3DSFlow: true,
				},
			},
			Phone:      "+333333333",
			PostalCode: "77777",
		},
		PlanID:    "PLAN ID",
		StartDate: time.Now().String(),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscription ID: ", resp.Subscription.ID)
}
