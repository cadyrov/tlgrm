package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Bot(confByte []byte) {

	config, err := Parse(confByte)
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = config.Timeout
	users := make(map[string]*Config, 0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		userCnf, ok := users[update.Message.From.UserName]
		if !ok {
			userCnf = config.Copy()
			users[update.Message.From.UserName] = userCnf
		}

		re, cn := userCnf.Answer(update.Message.Text)

		reply := fmt.Sprintf(re)
		ms := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(ms)
		result := fmt.Sprintf("%v", userCnf.answers)
		if cn != nil {
			z := cn
			sendToMan(*z, update.Message.From.UserName, result, bot)
		}
	}
}

func sendToMan(channel int64, user string, message string, bot *tgbotapi.BotAPI) {
	reply := "@" + fmt.Sprintf(user + " " +message)
	msg := tgbotapi.NewMessage(channel,reply)
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println(user, err.Error())
	}

}
