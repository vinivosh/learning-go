package main

import (
	"fmt"
	"testing"
)

func TestBitcoin(t *testing.T) {
	t.Run("should stringify Bitcoin", func(t *testing.T) {
		btc := Bitcoin(10)
		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		wallet.Withdraw(Bitcoin(20))

		want := Bitcoin(30)
		assertBalance(t, wallet, want)
	})

}
