package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	_, error := acc.Withdraw(1000)

	balanceError := error.(BalanceError)

	if balanceError.reason != "Low account Balance" {
		t.Error("Account Balance is Low")
	}

}
