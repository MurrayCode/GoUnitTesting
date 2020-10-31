package Wallet

import (
	"fmt"
	"errors"
)

type Wallet struct{
	balance Bitcoin
}

type Stringer interface{
	String() string
}

type Bitcoin int
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func(w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func(w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}