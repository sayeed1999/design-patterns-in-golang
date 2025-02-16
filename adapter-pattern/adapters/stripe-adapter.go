package adapters

import "github.com/sayeed1999/design-patterns-in-golang/adapter-pattern/third-party/stripe"

// Implements PaymentProcessor, but takes Stripe 3rd-party lib via aggregation!

type StripeAdapter struct {
	StripeAPI *stripe.Stripe
}

func (s *StripeAdapter) MakePayment(amount float64, currency string) string {
	// terminating payment to all other currencies, because stripe doesn't support it
	if currency != "USD" {
		return "[Failed] Cannot initiate payment with Stripe!"
	}

	// calling stripe's api in the way stripe understands...
	return s.StripeAPI.Pay(amount)
}
