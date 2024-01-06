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

	resp, err := client.DeleteWalletCustomerPayment("123123", &rozetkapay.DeleteWalletCustomerSchema{
		OptionID: "1332",
		Type:     rozetkapay.PaymentMethodTypeWallet,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted: ", resp.Delete)
}
