package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.getBalance()
		expected := Bitcoin(10)
		if got != expected {
			t.Errorf("Got %s expected %s", got, expected)
		}
	})

	t.Run("Test Multiple Deposits", func(t *testing.T) {
		//adding to it
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(15))
		got := wallet.getBalance()
		expected := Bitcoin(25)
		if got != expected {
			t.Errorf("Got %s expected %s", got, expected)
		}
	})

	t.Run("Test Withdraw with balance", func(t *testing.T) {
		//adding to it
		wallet := Wallet{Bitcoin(10)}
		//wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(3))
		got := wallet.getBalance()
		expected := Bitcoin(7)
		if err != nil {
			t.Errorf("Got an error %s", err)
		}
		if got != expected {
			t.Errorf("Got %s expected %s", got, expected)
		}
	})

	t.Run("Test Withdraw with no balance", func(t *testing.T) {
		// adding to it
		wallet := Wallet{Bitcoin(3)}
		err := wallet.Withdraw(Bitcoin(5))

		if err == nil {
			t.Fatal("Wanted an error didnt got one")
		}
		got := wallet.getBalance()
		expected := Bitcoin(3) //balance didnt decreased
		if got != expected {
			t.Errorf("Got %s expected %s", got, expected)
		}

	})
}
