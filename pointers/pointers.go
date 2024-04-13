package pointers

import "fmt"

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

func (wallet *Wallet) Withdraw(withdrawal Bitcoin) {
  if withdrawal > wallet.balance {
    return 
  }

  wallet.balance -= withdrawal
}

func (wallet *Wallet) Balance() Bitcoin {
  return wallet.balance
}
