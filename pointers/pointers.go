package pointers

import (
	"fmt"
)

//ErrNoFund because no fund
// var ErrNoFund = errors.New("not enough money in your wallet")

//ErrNoFund because no fund
type ErrNoFund struct {
	msg     string
	balance Bitcoin
}

func (e *ErrNoFund) Error() string {
	return fmt.Sprintf("%s , Current balance is %s", e.msg, e.balance)
}

//Bitcoin is a new currency
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet to store money
type Wallet struct {
	ballance Bitcoin
}

//Deposit into the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.ballance += amount
}

//GetBalance of the wallet
func (w *Wallet) getBalance() Bitcoin {
	return w.ballance
}

//Withdraw money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.ballance {
		return &ErrNoFund{"Not enough fund", w.ballance}
	}
	w.ballance -= amount
	return nil
}
