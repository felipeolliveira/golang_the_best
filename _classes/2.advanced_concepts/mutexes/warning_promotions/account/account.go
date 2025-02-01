package account

import "sync"

type AccountWithPromotion struct {
	sync.Mutex // Promotion
	balance    int
}

func (a *AccountWithPromotion) Deposit(amount int) {
	a.Lock()
	defer a.Unlock()

	a.balance += amount
}

type AccountPrivate struct {
	mu      sync.Mutex // Private
	balance int
}

func (a *AccountPrivate) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}
