package telegram

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/RACE-Game/ton-deposit/internal/domain/notification"
)

func (t *Telegram) SendNotificationToUser(ctx context.Context, userID int64, notification notification.Notification) error {
	simpleMsg := tgbotapi.NewMessage(userID, notification.Text)
	if notification.Button != nil {
		simpleMsg.ReplyMarkup = NewReplyMarkup(notification.Button.Text, notification.Button.URL)
	}

	_, err := t.tgClient.Send(simpleMsg)
	if err != nil {
		return fmt.Errorf("unable to send message: %w", err)
	}

	/*
	 sendMessage, response: {"ok":false,"error_code":403,"description":"Forbidden: bot was blocked by the user"}

	*/

	return nil
}

func NewReplyMarkup(text, link string) tgbotapi.InlineKeyboardMarkup {
	keyboardRows := [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonWebApp(text, tgbotapi.WebAppInfo{URL: link}),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonURL("Papa", "https://google.com"),
		// 	tgbotapi.NewInlineKeyboardButtonData("Mama", "AM AM I said!!!"),
		// ),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Fixic", "FuckIt"),
		// ),

	}

	startKeys := tgbotapi.NewInlineKeyboardMarkup(
		keyboardRows...,
	)

	return startKeys
}
