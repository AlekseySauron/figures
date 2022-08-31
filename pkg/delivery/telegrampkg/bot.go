package telegrampkg

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func Send(chat_id string, res string) {

	token := viper.GetString("telegram.token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}

	chat_id_int, _ := strconv.Atoi(chat_id)
	bot.Send(tgbotapi.NewMessage(int64(chat_id_int), fmt.Sprint(res)))

}
