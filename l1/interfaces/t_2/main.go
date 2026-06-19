package main

import (
	pp "PaymentProcessor"
	"fmt"
)

func main() {
	tbank, tBankKey := pp.NewTBank(154.12)
	sberbank := pp.NewSberbank("test", "123", 100.64)
	alphabank := pp.NewAlphaBank("SOME_ULTRA_SECRET_KEY", 346.28)

	banks := [3]pp.PaymentProcessor{
		&sberbank,
		&alphabank,
		&tbank,
	}

	errCount := len(banks)
	sberbank.LogIn("test", "123")
	go tbank.SendKey(tBankKey)

	errs := make(chan error, errCount)
	for _, bank := range banks {
		go func() {
			errs <- bank.ProcessPayment(100.0)
		}()
	}

	<-tbank.Unlocked

	for err := range errs {
		if err != nil {
			fmt.Println(err.Error())
		}
		errCount -= 1
		if errCount == 0 {
			close(errs)
		}
	}
}
