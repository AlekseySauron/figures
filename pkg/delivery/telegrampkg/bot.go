package telegrampkg

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

var bot *tgbotapi.BotAPI

func initBot() {
	var err error

	token := viper.GetString("telegram.token")
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err, " token", token)
		return
	}
}

func Send(chat_id string, res string) {

	chat_id_int, err := strconv.Atoi(chat_id)
	if err != nil || chat_id_int <= 0 {
		log.Fatal("Ошибка формата chat_id:", err)
		return
	}

	if bot == nil {
		initBot()
		// log.Fatal("Ошибка Бот не создан:", err)
		// return
	}

	bot.Send(tgbotapi.NewMessage(int64(chat_id_int), fmt.Sprint(res)))

}
