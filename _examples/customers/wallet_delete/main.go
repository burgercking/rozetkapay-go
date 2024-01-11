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

	resp, err := manager.DeleteWalletCustomerPayment("123123", &rozetkapay.DeleteWalletCustomerSchema{
		OptionID: "1332",
		Type:     rozetkapay.PaymentMethodTypeWallet,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted: ", resp.Delete)
}
