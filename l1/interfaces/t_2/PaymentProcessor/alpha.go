package paymentprocessor

import "math/rand"

type AlphaBank struct {
	ApiKey  string
	balance float64
}

const correctApiKey string = "SOME_ULTRA_SECRET_KEY"

func (a AlphaBank) checkApiKey() bool {
	return a.ApiKey == correctApiKey
}

func NewAlphaBank(apiKey string, balance float64) AlphaBank {
	return AlphaBank{
		ApiKey:  apiKey,
		balance: balance,
	}
}

func (a *AlphaBank) ProcessPayment(ammount float64) error {
	if !a.checkApiKey() {
		return ErrProviderUnavailable
	}
	if a.balance-ammount < 0 {
		return ErrInvalidAmount
	}

	if rand.Intn(4) == 0 { // about 75% of the time
		return ErrProviderUnavailable
	}

	a.balance -= ammount

	return nil
}

func (a AlphaBank) GetBalance() float64 {
	return a.balance
}
