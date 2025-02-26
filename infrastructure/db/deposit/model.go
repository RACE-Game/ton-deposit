package deposit

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Deposite struct {
	ID          int64            `db:"id"`
	Token       *string          `db:"token"`
	UserID      int64            `db:"user_id"`
	Amount      int64            `db:"amount"`
	TxHash      *string          `db:"tx_hash"`
	ConfirmedAt pgtype.Timestamp `db:"confirmed_at"`
	Wallet      string           `db:"wallet"`
	CreatedAt   pgtype.Timestamp `db:"created_at"`
}

type Replenishment struct {
	ID          int64     `db:"id"`
	Token       string    `db:"token"`
	Decimals    int16     `db:"decimals"`
	ClaimID     int64     `db:"claim_id"`
	UserID      int64     `db:"user_id"`
	Wallet      string    `db:"wallet"`
	Amount      int64     `db:"amount"`
	TXHash      string    `db:"tx_hash"`
	TXLT        int64     `db:"tx_lt"`
	TXTimestamp time.Time `db:"tx_timestamp"`
	CreatedAt   time.Time `db:"created_at"`
}
