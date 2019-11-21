package main

import (
	bot "./bot"
	"io/ioutil"
	"log"
)

func main() {
	if confByte, err := ioutil.ReadFile("./resources/ssupperbot/conf.yaml"); err == nil {
		bot.Bot(confByte)
	} else {
		log.Fatalln(err)
	}
}

