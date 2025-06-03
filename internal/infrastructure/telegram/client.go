package telegramClient

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"log"
	"time"
)

const OnText = telebot.OnText

type TelegramClient interface {
	Start()
	SendMessage(chatID int64, text string) error
	Stop()
	RegisterHandler(endpoint interface{}, handler telebot.HandlerFunc)
}

type telegramBot struct {
	bot   *telebot.Bot
	token string
}

func NewTelegramBot(token string, timeout time.Duration) (TelegramClient, error) {
	settings := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: timeout,
		},
	}
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}

	return &telegramBot{
		bot:   bot,
		token: token,
	}, nil

}
func (t *telegramBot) Start() {
	log.Printf("Starting Telegram bot with token: %s", t.token)
	t.bot.Start()
}

func (t *telegramBot) Stop() {
	log.Printf("Stopping Telegram bot.")
	t.bot.Stop()
}

func (t *telegramBot) SendMessage(chatID int64, text string) error {
	recipient := &telebot.Chat{
		ID: chatID,
	}

	_, err := t.bot.Send(recipient, text)
	if err != nil {
		log.Printf("Error sending message to chat %d: %v", chatID, err)
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (t *telegramBot) RegisterHandler(endpoint interface{}, handler telebot.HandlerFunc) {
	t.bot.Handle(endpoint, handler)
}
