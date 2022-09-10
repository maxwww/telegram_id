package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	token := os.Getenv("TOKEN")

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot := NewBot(botApi)

	log.Println(bot.Start())
}

type Bot struct {
	BotAPI *tgbotapi.BotAPI
}

func NewBot(botAPI *tgbotapi.BotAPI) *Bot {
	return &Bot{
		BotAPI: botAPI,
	}
}

func (bot *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.BotAPI.GetUpdatesChan(u)

	for update := range updates {
		go bot.handleUpdate(update)
	}

	return nil
}

func (bot *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	var chatId = update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Your id is %d", update.Message.From.ID))
	_, err := bot.BotAPI.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
