package transactions

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Amount    int       `db:"amount"`
	From      uuid.UUID `db:"sender_wallet"`
	To        uuid.UUID `db:"receiver_wallet"`
	CreatedAt time.Time `db:"transfer_date"`
}
