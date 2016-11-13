package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tucnak/telebot"
)

type QA struct {
	Q, A string
}

// copied verbatim from http://stackoverflow.com/a/31129967
func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	bot, err := telebot.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan telebot.Message, 1000)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		// fmt.Printf("%+v", message)
		if message.Text == "/ping" {

			bot.SendMessage(message.Chat, "PONG "+message.Sender.FirstName+"!", nil)
		}
		if message.Text == "/test" {
			bot.SendMessage(message.Chat, "kina test, "+message.Sender.FirstName+"?", nil)
		}
		if message.Text == "/rand" {
			qa := QA{}
			getJson(os.Getenv("API_HOST_ADDR"), &qa)
			bot.SendMessage(message.Chat, qa.Q+": "+qa.A, nil)
		}
		if message.Text == "/help" {
			const helpOpts = `
        available options:
              /help - shows this menu
              /ping - get a pong
              /rand - get random fun.tra words`

			bot.SendMessage(message.Chat, helpOpts, nil)
		}

	}
}
