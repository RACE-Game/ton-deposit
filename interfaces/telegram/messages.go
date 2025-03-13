package telegram

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/RACE-Game/ton-deposit/assets"
	"github.com/RACE-Game/ton-deposit/internal/domain/telegram"
)

const (
	inviteMessage          = "Provides 500 $MEMEPOLIS for invitee, 1000 $MEMEPOLIS for inviter, extra drops & score multipliers."
	communityLinkTrackable = "https://t.me/+LHprPFE2_FwwODEy"
)

var howToPlayCallbackData = "howtoplay"

var communityLink string = communityLinkTrackable

func (t *Telegram) messages(ctx context.Context, msg *tgbotapi.Message) error {
	if msg.WebAppData != nil {
		t.logger.Infof("WebAppData: %v", msg.WebAppData)
		//answerWebAppQuery

	}
	msgText := msg.Text

	refLink := t.userController.GetReferalLink(ctx, msg.From.ID)
	inviteReferal := " Mematrix Alpha Invite - " + refLink + " " + t.inviteMessage

	keyboardRows := [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonWebApp("Play", tgbotapi.WebAppInfo{URL: t.miniAppURL}),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonURL("Papa", "https://google.com"),
		// 	tgbotapi.NewInlineKeyboardButtonData("Mama", "AM AM I said!!!"),
		// ),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Fixic", "FuckIt"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "Join community",
				URL:  &t.communityLink,
				//SwitchInlineQuery: &inviteReferal,
			},
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "Invite",
				//	URL: ,
				SwitchInlineQuery: &inviteReferal,
			},
		),
	}

	if t.howToPlay != "" {
		keyboardRows = append(keyboardRows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.InlineKeyboardButton{
					Text:         "How to play",
					CallbackData: &howToPlayCallbackData,
				},
			),
		)
	}

	startKeys := tgbotapi.NewInlineKeyboardMarkup(
		keyboardRows...,
	)

	// wai := tgbotapi.WebAppInfo{
	// 	URL: t.miniAppURL,
	// }

	// startGameBtn := tgbotapi.NewKeyboardButtonWebApp("PLAY", wai)
	// startGameKeyboard := tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(startGameBtn),
	// )

	switch {
	case msgText == "/start":
		// msgReply := tgbotapi.NewMessage(msg.Chat.ID, "") //"Welcome to MeMePolis"
		// msgReply.ReplyMarkup = startKeys
		startPicture, _ := assets.Assets.ReadFile(t.startPicture)
		startImg := tgbotapi.FileBytes{"meme", startPicture}
		photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, startImg)
		photoMsg.ReplyMarkup = startKeys

		_, err := t.tgClient.Send(photoMsg)
		if err != nil {
			return fmt.Errorf("unable to send message: %w", err)
		}

		//tgbotapi.NewEditMessageReplyMarkup(msg.Chat.ID, msg.MessageID, startKeys)

		tgUser := telegram.User{
			TelegramID:       msg.From.ID,
			TelegramUserName: msg.From.UserName,
			FirstName:        msg.From.FirstName,
			LastName:         msg.From.LastName,
			LanguageCode:     msg.From.LanguageCode,
			IsPremium:        msg.From.IsPremium,
			ChatID:           msg.Chat.ID,
			StartMessageID:   int64(msg.MessageID),
			LastMessageAt:    msg.Time(),
		}

		exist, err := t.userController.UserExist(ctx, tgUser.TelegramID)
		if err != nil {
			return fmt.Errorf("unable to check user: %w", err)
		}

		if !exist {
			err = t.userController.SaveUser(ctx, tgUser)
			if err != nil {
				return fmt.Errorf("unable to save user: %w", err)
			}
		} else {
			err = t.userController.UpdateUser(ctx, tgUser)
			if err != nil {
				return fmt.Errorf("unable to update user: %w", err)
			}
		}

		return nil

	case referalStart.MatchString(msgText):
		startMessageArray := strings.Split(msgText, " ")
		if len(startMessageArray) != 2 {
			return fmt.Errorf("unable parse referrerUserID from msg: %s", msgText)
		}
		referrerUserID, err := strconv.ParseInt(startMessageArray[1], 10, 64)
		if err != nil {
			return fmt.Errorf("unable parse referrerUserID from msg: %s %w", msgText, err)
		}

		referalUserID := msg.From.ID

		err = t.userController.UserJoined(ctx, referrerUserID, referalUserID)
		if err != nil {
			return fmt.Errorf("unable to save referal: %w", err)
		}

		startPicture, _ := assets.Assets.ReadFile(t.startPicture)

		startImg := tgbotapi.FileBytes{"meme", startPicture}
		photoMsg := tgbotapi.NewPhoto(msg.Chat.ID, startImg)
		photoMsg.ReplyMarkup = startKeys

		_, err = t.tgClient.Send(photoMsg)
		if err != nil {
			return fmt.Errorf("unable to send message: %w", err)
		}

		tgUser := telegram.User{
			TelegramID:       msg.From.ID,
			TelegramUserName: msg.From.UserName,
			FirstName:        msg.From.FirstName,
			LastName:         msg.From.LastName,
			LanguageCode:     msg.From.LanguageCode,
			IsPremium:        msg.From.IsPremium,
			ChatID:           msg.Chat.ID,
			StartMessageID:   int64(msg.MessageID),
			LastMessageAt:    msg.Time(),
		}

		exist, err := t.userController.UserExist(ctx, tgUser.TelegramID)
		if err != nil {
			return fmt.Errorf("unable to check user: %w", err)
		}

		if !exist {
			err = t.userController.SaveUser(ctx, tgUser)
			if err != nil {
				return fmt.Errorf("unable to save user: %w", err)
			}
		} else {
			err = t.userController.UpdateUser(ctx, tgUser)
			if err != nil {
				return fmt.Errorf("unable to update user: %w", err)
			}
		}

		return nil

	case msgText == "invite":

		return nil

	case msgText == "test1":
		// btn := tgbotapi.KeyboardButton{
		// 	Text:           "RRRRR conTAAACT!!!",
		// 	RequestContact: true,
		// }

		//urlbtn := "sdfsdfsd"

		return nil
	}

	return nil
}
