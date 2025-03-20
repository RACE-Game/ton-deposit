package withdrawal

import "time"

type Withdrawal struct {
	Amount int64
	TS     time.Time
	TXID   string
}
