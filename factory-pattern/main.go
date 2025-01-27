package main

import (
	"fmt"

	"github.com/sayeed1999/design-patterns-in-golang/factory-pattern/payment"
)

func main() {
	paymentTypes := []string{"debitcard", "paypal", "banktransfer", "crypto"}
	amount := 100.0

	for _, paymentType := range paymentTypes {
		// Get the payment processor
		payProcessor, err := payment.PaymentFactory(paymentType)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Process the payment
		err = payProcessor.Pay(amount)
		if err != nil {
			fmt.Println("Payment failed:", err)
		}
	}
}

// Output:
// Paid 100.00 using Debit Card
// Paid 100.00 using PayPal
// Paid 100.00 using Bank Transfer
// Error: unsupported payment type: crypto

// Task: Enable crypto payment!
