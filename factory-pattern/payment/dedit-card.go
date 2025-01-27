package payment

import "fmt"

// DebitCard represents payment via debit card

type deditCard struct{}

func (c *deditCard) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Debit Card\n", amount)
	return nil
}
