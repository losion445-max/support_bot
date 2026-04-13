package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	bot, err := tele.NewBot(tele.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		log.Panic(err)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Поддержка")
	})

	bot.Start()
}
