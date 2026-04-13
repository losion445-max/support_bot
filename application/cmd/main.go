package main

import (
	"log"
	"strconv"
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

	bot.Handle(tele.OnText, func(c tele.Context) error {
		senderID := c.Sender().ID

		adminID, err := strconv.ParseInt(cfg.AdminID, 10, 64)
		if err != nil {
			log.Panicf("Критическая ошибка: AdminID в конфиге не является числом: %v", err)
		}

		if senderID == adminID {
			if c.Message().ReplyTo == nil {
				return c.Send("Нажми Reply на сообщение пользователя чтобы ответить")
			}

			orig := c.Message().ReplyTo.OriginalSender
			if orig == nil {
				log.Println(err)
				return c.Send("Пользотель использует настрокий приватности")
			}

			_, err := c.Bot().Send(&tele.User{ID: orig.ID}, c.Text())
			if err != nil {
				log.Println(err)
			}

			return c.Send("Сообщение доставлено")
		}

		_, err = c.Bot().Forward(&tele.User{ID: adminID}, c.Message())
		if err != nil {
			log.Println(err)
		}

		return c.Send("Ваше сообщение получено")
	})

	bot.Start()
}
