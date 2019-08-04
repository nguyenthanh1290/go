package account

import (
	"sync"
)

// Account represents a bank account.
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open opens an account with the initial deposit. A negative deposit returns nil.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit, closed: false}
}

// Close will close the account and return the current balance.
func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if !acc.closed {
		acc.closed = true
		payout = acc.balance
		ok = true
	}
	return
}

// Balance will return account's current balance, or (0, false) if the account is closed.
func (acc *Account) Balance() (balance int64, ok bool) {
	acc.RLock()
	defer acc.RUnlock()

	if acc.closed {
		return
	}
	return acc.balance, true
}

// Deposit will deposit or withdraw (if negative) the specified amount.
// Withdrawals will not succeed if they result in a negative balance.
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return
	}

	newBalance = acc.balance + amount
	if newBalance < 0 {
		return
	}
	acc.balance = newBalance
	ok = true
	return
}
