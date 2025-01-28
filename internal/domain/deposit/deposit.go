package deposit

import "time"

type Deposit struct {
	Amount     int64
	TXID       string
	TelegramID int64
	TS         time.Time
}
