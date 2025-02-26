package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) IsMember(ctx context.Context, userID, chatID int64) (bool, error) {
	chatMember, err := t.tgClient.GetChatMember(
		tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
				ChatID:             chatID,
				SuperGroupUsername: "",
				UserID:             userID,
			},
		},
	)
	if err != nil {
		return false, err
	}

	switch chatMember.Status {
	case "":
		return false, nil
	case "left":
		return false, nil
	case "kicked":
		return false, nil
	}

	return true, nil
}
