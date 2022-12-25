package account

// TDD Bank Account app
import (
	"fmt"
)

type Account struct {
	balance float64
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

type BalanceError struct {
	amount     float64
	accountBal float64
	reason     string
}

func (b BalanceError) Error() string {
	return fmt.Sprintf(" You cannot withdraw %f due to  %s from %f", b.amount, b.reason, b.accountBal)
}
func (acc *Account) Withdraw(amount float64) (float64, error) {

	remainigBalance := acc.balance - amount
	if amount > acc.balance {
		return remainigBalance, BalanceError{1000, remainigBalance, "Low account Balance"}
	}
	return remainigBalance, nil
}
