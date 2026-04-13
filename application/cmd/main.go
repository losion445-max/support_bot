package main

import (
	"log"
	"time"

	"github.com/losion445-max/support_bot-golang/internal/config"
	tele "gopkg.in/telebot.v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := tele.NewBot(tele.Settings{
		Token: cfg.Token,
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
