package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5741530807:AAEemmt_lmDIv9kXEWwOpb_GBprO_buk2rk")
	if err != nil {
		fmt.Print(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		returnmsg := "asdasd"
		if update.Message.Text == "naber" {
			returnmsg = "iyi sen"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, returnmsg)
		bot.Send(msg)
	}
}
