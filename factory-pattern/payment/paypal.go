package payment

import "fmt"

// PayPal represents payment via PayPal

type payPal struct{}

func (p *payPal) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
	return nil
}
