package main

import (
	"fmt"

	"github.com/sayeed1999/design-patterns-in-golang/adapter-pattern/adapters"
	"github.com/sayeed1999/design-patterns-in-golang/adapter-pattern/third-party/paypal"
	"github.com/sayeed1999/design-patterns-in-golang/adapter-pattern/third-party/stripe"
)

func main() {
	stripeAdapter := &adapters.StripeAdapter{
		StripeAPI: &stripe.Stripe{},
	}

	fmt.Println(stripeAdapter.MakePayment(100.64, "BDT"))
	fmt.Println(stripeAdapter.MakePayment(100.64, "USD"))

	paypalAdapter := &adapters.PaypalAdapter{
		PaypalAPI: &paypal.Paypal{},
	}

	fmt.Println(paypalAdapter.MakePayment(100.64, "BDT"))
}
