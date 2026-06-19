package paymentprocessor

import (
	"errors"
	"math/rand"
)

type Sberbank struct {
	login      string
	password   string
	isLoggedIn bool
	balance    int // храним копейки как целое число
}

func NewSberbank(login string, password string, balance float64) Sberbank {
	return Sberbank{
		login:      login,
		password:   password,
		isLoggedIn: false,
		balance:    int(balance * 100),
	}
}

func (s *Sberbank) LogIn(login string, password string) error {
	if s.login != login {
		return errors.New("login does not exists")
	}
	if s.password != password {
		return errors.New("password is wrong")
	}

	s.isLoggedIn = true
	return nil
}

func (s *Sberbank) ProcessPayment(ammount float64) error {
	if !s.isLoggedIn {
		return ErrProviderUnavailable
	}
	if ammount*100 > float64(s.balance) {
		return ErrInvalidAmount
	}

	if rand.Intn(4) == 0 { // about 75% of the time
		return ErrProviderUnavailable
	}

	s.balance -= int(ammount * 100)
	return nil
}

func (s Sberbank) GetBalance() float64 {
	return float64(s.balance) / 100
}
