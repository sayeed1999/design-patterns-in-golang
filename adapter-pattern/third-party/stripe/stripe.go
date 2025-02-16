package stripe

import "fmt"

// Considering stripe is a 3rd party provider that doesn't align with our PaymentProcessor interface..

type Stripe struct{}

// Notice the problem! Stripe only works for USD!!
func (s *Stripe) Pay(amount float64) string {
	return fmt.Sprintf("[Success] Paid %.2f USD via Stripe!", amount)
}
