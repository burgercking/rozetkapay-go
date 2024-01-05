package main

import (
	"log"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
)

func main() {
	var (
		cfg = rozetkapay.NewConfig(rozetkapay.DevEnvironmentProjectID, rozetkapay.DevEnvironmentPassword).
			SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	resp, err := client.DeleteWalletCustomerPayment("123123", &rozetkapay.DeleteWalletCustomerParams{
		OptionID: "1332",
		Type:     rozetkapay.PaymentMethodTypeWallet,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted: ", resp.Delete)
}
