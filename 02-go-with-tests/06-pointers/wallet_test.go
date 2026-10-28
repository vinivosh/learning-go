package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := 10
	wallet.Deposit(want)

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
