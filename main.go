package main

import (
	"io/ioutil"
	"log"
	"telegramm/bot"
)

func main() {
	if configByte, err := ioutil.ReadFile("./resources/ssupperbot/conf.yaml"); err == nil {
		bot.Bot(configByte)
	} else {
		log.Fatalln(err)
	}
}
