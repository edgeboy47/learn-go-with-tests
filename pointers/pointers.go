package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient Funds")

type Bitcoin int

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
  balance Bitcoin
}

func (wallet *Wallet) Deposit(deposit Bitcoin) {
  wallet.balance += deposit
}

func (wallet *Wallet) Withdraw(withdrawal Bitcoin) error {
  if withdrawal > wallet.balance {
    return ErrInsufficientFunds
  }

  wallet.balance -= withdrawal
  return nil
}

func (wallet *Wallet) Balance() Bitcoin {
  return wallet.balance
}
