package payment

// Payment is the interface for all payment methods
type Payment interface {
	Pay(amount float64) error
}
