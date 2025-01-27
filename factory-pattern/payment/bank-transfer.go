package payment

import "fmt"

// BankTransfer represents payment via bank transfer

type bankTransfer struct{}

func (b *bankTransfer) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Bank Transfer\n", amount)
	return nil
}
