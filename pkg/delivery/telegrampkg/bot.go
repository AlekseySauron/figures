package telegrampkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func Send(chatID int64, msg string) {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("telegram.token"))
	if err != nil {
		return
	}
	_, _ = bot.Send(tgbotapi.NewMessage(chatID, msg))
}
