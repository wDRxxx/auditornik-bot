package main

import (
	"github.com/joho/godotenv"
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

var app config.AppConfig

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	bot := Bot()
	app.Bot = bot

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	routes(&app)

	log.Println("Bot started!")
	bot.Start()
}

func Bot() *telebot.Bot {
	token := mustToken()

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func mustToken() string {
	return os.Getenv("TG_TOKEN")
}
