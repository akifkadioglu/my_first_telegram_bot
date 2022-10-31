package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
	fmt.Println(os.Getenv("TELEGRAM_API"))
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API"))
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
