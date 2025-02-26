package telegram

import (
	"context"
	"log"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/RACE-Game/ton-deposit-service/assets"
)

type Telegram struct {
	tgClient       *tgbotapi.BotAPI
	userController UserController
	miniAppURL     string
	communityLink  string
	inviteMessage  string
	startPicture   string
	howToPlay      string
	logger         Logger
}

var referalStart = regexp.MustCompile(`^/start \d+$`)

func New(
	tgBotAPI, miniAppURL, communityLink, inviteMessage, startPicture, howToPlay string,
	userController UserController, logger Logger, debug bool,
) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(tgBotAPI)
	if err != nil {
		return nil, err
	}
	if debug {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Telegram{
		tgClient:       bot,
		miniAppURL:     miniAppURL,
		communityLink:  communityLink,
		inviteMessage:  inviteMessage,
		startPicture:   startPicture,
		howToPlay:      howToPlay,
		userController: userController,
		logger:         logger,
	}, nil
}

func (t *Telegram) Start(ctx context.Context) error {
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := t.tgClient.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil {
				err := t.messages(ctx, update.Message)
				if err != nil {
					t.logger.Errorf("unable to proceed message: %v", err)
				}

				continue
			}

			if update.CallbackQuery != nil {
				// Respond to the callback query, telling Telegram to show the user
				// a message with the data received.
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := t.tgClient.Request(callback); err != nil {
					t.logger.Errorf("newcallback error: %s", err.Error())
				}

				if update.CallbackQuery.Data == howToPlayCallbackData {
					if t.howToPlay == "" {
						continue
					}
					err := t.SendHowToPlay(update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID)
					if err != nil {
						t.logger.Errorf("unable to send how to play: %v", err)
					}
					continue
				}

				// And finally, send a message containing the data received.
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				if _, err := t.tgClient.Send(msg); err != nil {
					t.logger.Errorf("newmessage error: %s", err.Error())
				}

				continue
			}

			// if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			// 	t.logger.Infof("Message raw: %+v", update.Message)
			// 	continue
			// }

			// If we got a message
			// t.logger.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// t.logger.Infof("Update raw: %+v", update)
			// t.logger.Infof("Message raw: %+v", update.Message)

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			//t.tgClient.Send(msg)

		}
	}()

	return nil
}

func (t *Telegram) SendInvite(ctx context.Context, userID int64, chatID int64) error {
	refLink := t.userController.GetReferalLink(ctx, userID)
	inviteMsg := tgbotapi.NewMessage(chatID, refLink)

	_, err := t.tgClient.Send(inviteMsg)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telegram) SendHowToPlay(userID int64, chatID int64) error {
	// Text to respond to "How to Play":
	var howToPlay = "Farm 6 legit tokens from the start 💰🚀\n More play = more tokens to farm 🎮➕🌱\n Why 1 lambo when you can get 6? 🏎️x6️⃣"

	howToplayTextMsg := tgbotapi.NewMessage(chatID, howToPlay)
	_, err := t.tgClient.Send(howToplayTextMsg)
	if err != nil {
		return err
	}

	howToPlayFile, _ := assets.Assets.ReadFile(t.howToPlay)
	howToPlayTGFile := tgbotapi.FileBytes{"howtoplay.gif", howToPlayFile}
	howToPlayMsg := tgbotapi.NewVideo(chatID, howToPlayTGFile)

	_, err = t.tgClient.Send(howToPlayMsg)
	if err != nil {
		return err
	}

	return nil
}
