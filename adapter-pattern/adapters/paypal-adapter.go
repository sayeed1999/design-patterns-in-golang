package adapters

import (
	"github.com/sayeed1999/design-patterns-in-golang/adapter-pattern/third-party/paypal"
)

// Implements PaymentProcessor, but takes Paypal 3rd-party lib via aggregation! (not composition)

type PaypalAdapter struct {
	PaypalAPI *paypal.Paypal
}

func (p *PaypalAdapter) MakePayment(amount float64, currency string) string {
	// payment amount converted to the format that paypal's api understands...
	return p.PaypalAPI.SendPayment(int(amount), currency)
}
