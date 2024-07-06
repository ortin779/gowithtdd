package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposite", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(30))

		fmt.Println(err)

		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()

	if wallet.Balance() != expected {
		t.Errorf("expected %q, but got %q", expected, wallet.Balance())
	}
}

func assertError(t testing.TB, expected, got error) {
	t.Helper()

	if got == nil {
		t.Fatalf("didn't get the error but expected one")
	}

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("didn't expect an error, but got one.")
	}
}
