package paymentprocessor

import "errors"

type PaymentProcessor interface {
	ProcessPayment(ammount float64) error
}

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
)
