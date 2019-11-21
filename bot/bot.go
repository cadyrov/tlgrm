package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func Bot(confByte []byte) {

	baseConfig, err := Parse(confByte)
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(baseConfig.Token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = baseConfig.Timeout
	updateChannels, err := bot.GetUpdatesChan(updateConfig)

	users := make(map[string]*Config, 0)

	for channel := range updateChannels {
		if channel.Message == nil {
			continue
		}
		userConfig, ok := users[channel.Message.From.UserName]
		if !ok {
			userConfig = baseConfig.Copy()
			users[channel.Message.From.UserName] = userConfig
		}

		nextMessage, employeeChannel := userConfig.Answer(channel.Message.Text)

		prepareMessage := tgbotapi.NewMessage(channel.Message.Chat.ID, nextMessage)
		if _, err = bot.Send(prepareMessage); err != nil {
			log.Println(err)
		}
		historyAnswers := fmt.Sprintf("%v", userConfig.answers)
		if employeeChannel != nil {
			sendToMan(*employeeChannel, channel.Message.From.UserName, historyAnswers, bot)
		}
	}
}

func sendToMan(channel int64, user string, message string, bot *tgbotapi.BotAPI) {
	reply := "@" + user + " " + message
	preparedMessage := tgbotapi.NewMessage(channel, reply)
	_, err := bot.Send(preparedMessage)
	if err != nil {
		log.Println(user, err.Error())
	}
}
