package main

import (
	"io"
	"log"
	"net/http"

	"github.com/burgercking/rozetkapay-go"
	"github.com/burgercking/rozetkapay-go/examples"
	"github.com/google/uuid"
)

func main() {
	var (
		cfg = rozetkapay.NewConfig(rozetkapay.DevEnvironmentProjectID, rozetkapay.DevEnvironmentPassword).
			SetCallbackURL(examples.DevEnvironmentCallbackURL)
		client = rozetkapay.NewClient(cfg)
	)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			log.Println("err1: ", err)
			return
		}

		callback, err := client.GetPaymentCallbackFromBytes(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("err3: ", err)
			return
		}

		switch callback.Details.StatusCode {
		case rozetkapay.StatusCodePaymentWasRefunded:
			log.Println("Payment refund: ", callback.ExternalID)
			log.Println("Status: ", callback.Details.Status)
		case rozetkapay.StatusCodeTransactionSuccessful:
			log.Println("Transaction: ", callback.ExternalID)
			log.Println("Status: ", callback.Details.Status)
		default:
			log.Println("ExternalID: ", callback.ExternalID)
			log.Println("Status: ", callback.Details.Status)
			log.Println("Status code: ", callback.Details.StatusCode)
		}

		w.WriteHeader(http.StatusOK)
	})

	payment, err := client.CreatePayment(&rozetkapay.CreatePaymentParams{
		Amount:      2000,
		Currency:    "UAH",
		ExternalID:  uuid.New().String(),
		Mode:        rozetkapay.PaymentModeHosted,
		CallbackURL: cfg.CallbackURL,
		Confirm:     false,
		Description: "Test payment",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Payment link:", payment.GetLink())
	log.Fatal(http.ListenAndServe(":5000", nil))
}
