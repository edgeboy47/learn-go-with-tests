package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

    assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

    assertBalance(t, wallet, 10)
	})
}
