package factory

/*
	Factory method creational design pattern allows creating objects without
	having to specify the exact type of the object that will be created.
*/

import (
	"fmt"
)

// PaymentMethod ...
type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	// Cash ...
	Cash = 1
	// DebitCard ...
	DebitCard = 2
)

// GetPaymentMethod ...
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

// CashPM ...
type CashPM struct {
}

// DebitCardPM ...
type DebitCardPM struct {
}

// Pay ...
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

// Pay ...
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
