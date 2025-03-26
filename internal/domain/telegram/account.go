package telegram

type Account struct {
	TelegramID       int64
	TelegramUserName string
	FirstName        string
	LastName         string

	SolanaAccount string
}
