package main

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBitcoinsEquals := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBitcoinsEquals(t, got, want)
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(5)
		if err != nil {
			t.Errorf("Ошибка, хотя не должно быть. %s", err)
		}

		got := wallet.Balance()
		want := Bitcoin(5)
		assertBitcoinsEquals(t, got, want)
	})

	t.Run("test withdraw with not enough money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(11)
		if !errors.Is(err, ErrNotEnoughMoney) {
			t.Errorf("Нет ошибки, хотя должно быть. %s", err)
		}
	})
}
