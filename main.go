package main

import (
	"fmt"
	"strconv"

	"example.com/m/config"
	"example.com/m/controllers"
	"github.com/VictorAvelar/mollie-api-go/mollie"
)

func main() {
	client := config.GetClient()
	var payment mollie.Payment
	var err error

	/* create */
	// lang := mollie.Locale("en_US")
	// payment, err = controllers.Create(
	// 	client,
	// 	mollie.Payment{
	// 		// required
	// 		Amount:      &mollie.Amount{Currency: "EUR", Value: "50.00"},
	// 		Description: "test #1",
	// 		RedirectURL: "https://www.example.org/order12/complete",

	// 		WebhookURL: "https://www.example.org/order12/webhook",
	// 		Locale:     &lang,
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(payment)

	/* get list */
	payments, err := controllers.GetAll(client, &mollie.ListPaymentOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for i, payment := range payments.Embedded.Payments {
		fmt.Println(
			strconv.Itoa(i+1)+".",
			payment.ID,
			payment.Description,
			payment.Amount,
			payment.Status,
		)
	}

	/* get one by id(string) */
	payment, err = controllers.GetOne(client, "tr_DoVsKtFMZm", &mollie.PaymentOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(payment)

	// /* update by id */
	// payment, err := controllers.Update(
	// 	client,
	// 	"tr_Ewi9BxRacd",
	// 	mollie.Payment{
	// 		// Amount:      &mollie.Amount{} - non-existent parameter for update,
	// 		Description: "test-order1",
	// 		// RedirectURL: "https://www.example.org/test-order#1/complete",
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(payment)

	// /* delete by id */
	// payment, err := controllers.Delete(client, "tr_Ewi9BxRacd")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(payment.IsCancellable)
}
