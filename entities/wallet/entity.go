package wallet

import "github.com/google/uuid"

type Wallet struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	balance int64
}

func New(ownerID uuid.UUID) *Wallet {
	return &Wallet{
		ID:      uuid.New(),
		OwnerID: ownerID,
		balance: 0,
	}
}

func (w *Wallet) Balance() *int64 {
	return &w.balance
}
