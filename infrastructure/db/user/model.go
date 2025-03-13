package user

import "time"

type User struct {
	TelegramID       int64     `db:"telegram_id"`
	TelegramUserName string    `db:"tg_user_name"`
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	LanguageCode     string    `db:"language_code"`
	IsPremium        bool      `db:"is_premium"`
	ChatID           int64     `db:"chat_id"`
	StartMessageID   int64     `db:"start_message_id"`
	LastMessageAt    time.Time `db:"last_message_at"`
}
