package main

import (
	"fmt"

	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	replicate "github.com/sausheong/goreplicate"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}
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

		model := replicate.NewModel("stability-ai", "stable-diffusion", os.Getenv("DALL_E_MINI_VERSION"))

		model.Input["prompt"] = update.Message.Text
		model.Input["num_outputs"] = 1
		client := replicate.NewClient(os.Getenv("DALL_E_MINI_API"), model)
		err := client.Create()
		if err != nil {
			fmt.Println("bir hata")
		}
		err = client.Get(client.Response.ID)
		if err != nil {
			fmt.Println("bir hata")
		}
		content := "Something went wrong, try something else!"

		if len(client.Response.Output) > 0 {
			content = client.Response.Output[0]
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, content)
		bot.Send(msg)
	}
}
