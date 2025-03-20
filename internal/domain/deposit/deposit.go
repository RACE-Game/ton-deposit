package deposit

import "time"

type Deposit struct {
	TxLT        int64     `json:"TxLT"`
	Token       string    `json:"Token"`
	Wallet      string    `json:"Wallet"`
	Comment     string    `json:"Comment"`
	Amount      int64     `json:"Amount"`
	TXHash      string    `json:"TxHash"`
	Payload     string    `json:"Payload"`
	TXTimestamp string    `json:"TxTimestamp"`
	TXComment   string    `json:"TxComment"`
	CreatedAt   time.Time `json:"CreatedAt"`
}
