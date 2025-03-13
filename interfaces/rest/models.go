package rest

type ArbitrageResponse struct {
	Amount        int64  `json:"amount"`
	TransactionID string `json:"transaction_id"`
}

func NewArbitrageResponse(amount int64, transactionID string) ArbitrageResponse {
	return ArbitrageResponse{
		Amount:        amount,
		TransactionID: transactionID,
	}
}

type NotificationRequest struct {
	Text   string `json:"text"`
	Button Button `json:"button"`
}

type Button struct {
	Text string `json:"text"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type WalletRequest struct {
	UserID int64  `json:"user_id"`
	Wallet string `json:"wallet"`
}

type UserDataRequest struct {
	Data string `json:"data"`
}

type ClaimRequest struct {
	Token  string `json:"token"`
	Amount int64  `json:"amount"`
	Wallet string `json:"wallet"`
}
type ClaimResponse struct {
	TX            string `json:"tx"`
	Payload       string `json:"payload"`
	PayloadBOC    []byte `json:"payload_boc"`
	PayloadSigned []byte `json:"payload_signed"`
	ID            int64  `json:"id"`
}
