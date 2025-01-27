package payment

import "fmt"

// PaymentFactory creates a Payment object based on the payment type
func PaymentFactory(paymentType string) (Payment, error) {
	switch paymentType {
	case "debitcard":
		return &deditCard{}, nil
	case "paypal":
		return &payPal{}, nil
	case "banktransfer":
		return &bankTransfer{}, nil
	default:
		return nil, fmt.Errorf("unsupported payment type: %s", paymentType)
	}
}
