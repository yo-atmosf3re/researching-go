package practice

import (
	"errors"
	"researching-go/pkg/logger"
)

func deductMoney(balance *int, amount int) error {
	if *balance-amount < 0 || *balance == 0 || amount <= 0 {
		return errors.New("operation rejected")
	}
	*balance -= amount
	logger.Ptc("operation successful")
	return nil
}

func RememberingFunc() {
	myWallet := 500

	err := deductMoney(&myWallet, 200)
	if err != nil {
		logger.Ptc("error: ", err)
	}
	deductMoney(&myWallet, 200)
	deductMoney(&myWallet, 200)
	logger.Ptc("balance of my wallet: ", myWallet)
}
