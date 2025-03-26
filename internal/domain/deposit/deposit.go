package deposit

import (
	"time"

	"github.com/google/uuid"
)

type Deposit struct {
	TxLT        int64     `json:"TxLT"`
	Token       string    `json:"Token"`
	Wallet      string    `json:"Wallet"`
	Comment     string    `json:"Comment"`
	Amount      uint64    `json:"Amount"`
	TXHash      string    `json:"TxHash"`
	Payload     string    `json:"Payload"`
	TXTimestamp string    `json:"TxTimestamp"`
	TXComment   string    `json:"TxComment"`
	CreatedAt   time.Time `json:"CreatedAt"`
}

type Order struct {
	ID        uuid.UUID `json:"ID"`
	Token     string    `json:"Token"`
	UserID    int64     `json:"UserID"`
	Amount    uint64    `json:"Amount"`
	Wallet    string    `json:"Wallet"`
	CreatedAt time.Time `json:"CreatedAt"`
}
