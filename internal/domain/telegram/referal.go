package telegram

import "time"

type Referal struct {
	ReferrerID int64
	ReferalID  int64
	CreatedAt  time.Time
}
