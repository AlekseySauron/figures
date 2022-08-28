package telegrampkg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewBot() *TelegramBot {
	bot, err := tgbotapi.NewBotAPI("1901733643:AAHlKkQJrCaKS1c1SZigHXq6t8CUXO7eeWs")
	if err != nil {
		return nil
	}

	return &TelegramBot{bot}
}

func (tgbot *TelegramBot) Send(res string) {

	msg := tgbotapi.NewMessage(421964311, fmt.Sprint(res))
	tgbot.bot.Send(msg)

}
