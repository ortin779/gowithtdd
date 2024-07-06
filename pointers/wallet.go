package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(balance Bitcoin) {
	w.balance += balance
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
