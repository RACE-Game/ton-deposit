package telegram

import "time"

type User struct {
	TelegramID       int64
	TelegramUserName string
	FirstName        string
	LastName         string
	LanguageCode     string
	IsPremium        bool
	ChatID           int64
	StartMessageID   int64
	LastMessageAt    time.Time
}
