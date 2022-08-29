package telegrampkg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type TelegramBot struct {
	bot    *tgbotapi.BotAPI
	ChatID int64
}

func NewBot() *TelegramBot {
	token := viper.GetString("telegram.token")
	// bot, err := tgbotapi.NewBotAPI("1901733643:AAHlKkQJrCaKS1c1SZigHXq6t8CUXO7eeWs")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil
	}

	var ChatID int64
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		ChatID = update.Message.Chat.ID
		break
	}

	return &TelegramBot{bot, ChatID}
}

func (tgbot *TelegramBot) Send(res string) {

	// chat_id_str := viper.GetString("telegram.chat_id")
	// chat_id_int, _ := strconv.Atoi(chat_id_str)
	// chat_id_int64 := int64(chat_id_int)
	chat_id_int64 := tgbot.ChatID

	// msg := tgbotapi.NewMessage(421964311, fmt.Sprint(res))
	msg := tgbotapi.NewMessage(chat_id_int64, fmt.Sprint(res))
	tgbot.bot.Send(msg)

}
