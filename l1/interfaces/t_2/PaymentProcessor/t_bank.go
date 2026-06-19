package paymentprocessor

import (
	"errors"
	"math/rand"
)

type TBank struct {
	keyChan  chan int
	Unlocked chan bool
	balance  float64
	lock     int
	hasKey   bool
}

var primes [10]int = [10]int{2, 3, 5, 7, 11, 13, 17, 23, 29, 31}

const bigPrime = 7919

func NewTBank(startBalance float64) (TBank, int) {
	key := rand.Intn(10)

	return TBank{
		balance:  startBalance,
		lock:     bigPrime * key,
		keyChan:  make(chan int),
		hasKey:   false,
		Unlocked: make(chan bool),
	}, key
}

func (t *TBank) SendKey(key int) error {
	if t.hasKey {
		return errors.New("there already is a key stored")
	}

	t.keyChan <- key
	t.hasKey = true

	return nil
}

// Я тут хотел сделать проверку ключа при проведении Processpayment, а не при создании структуры
// Так как в интерфейсе можно принимать только один параметр f64, то я решил сделать прием ключа
// через канал.
// Технически - эта реализация ломает идею интерфейса, и это плохой паттерн, но да ладно
func (t *TBank) ProcessPayment(ammount float64) error {
	if rand.Intn(4) == 0 { // about 75% of the time
		t.Unlocked <- false
		return ErrProviderUnavailable
	}

	key := <-t.keyChan
	t.hasKey = false
	if t.lock%key != 0 {
		t.Unlocked <- false
		return ErrProviderUnavailable
	}

	if t.balance-ammount < 0 {
		t.Unlocked <- true
		return ErrInvalidAmount
	}

	t.balance -= ammount
	t.Unlocked <- true
	return nil
}

func (t TBank) GetBalance() float64 {
	return t.balance
}
