package payment

// All payment methods should implement this common interface in the system!
type PaymentProcessor interface {
	MakePayment(amount float64, currency string) string
}
