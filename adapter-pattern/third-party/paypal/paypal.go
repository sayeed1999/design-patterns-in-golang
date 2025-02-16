package paypal

import "fmt"

// Considering paypal is a 3rd party provider that doesn't align with our PaymentProcessor interface..

type Paypal struct{}

// Paypal slightly aligns with our interface, but not supports fractional payment..
func (s *Paypal) SendPayment(amount int, currency string) string {
	return fmt.Sprintf("[Success] Paid %d %s via Paypal!", amount, currency)
}
