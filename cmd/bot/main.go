package main

import (
	"github.com/joho/godotenv"
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
	"github.com/wDRxxx/auditornik-bot/internal/parser"
	"github.com/wDRxxx/auditornik-bot/internal/storage"
	job_ticker "github.com/wDRxxx/auditornik-bot/pkg/job-ticker"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"sync"
	"time"
)

var app config.AppConfig
var wg sync.WaitGroup

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	bot, strg, prsr := run()
	app.Bot = bot

	if err != nil {
		log.Fatalf("error while getting storage: %s", err.Error())
	}

	repo := handlers.NewRepository(&app, strg, prsr)
	handlers.NewHandlers(repo)

	routes(&app)

	jt := job_ticker.New(handlers.Repo.Mailing)
	app.Ticker = jt

	wg.Add(1)
	go bot.Start()
	log.Println("Bot started!")

	// запускает ежедневную рассылку аудиторника
	wg.Add(1)
	go app.Ticker.Run()

	// wg тут как костыль, посзволяющий программе не прекращать свое выполнение
	wg.Wait()
}

// run создает объект бота и хранилища
func run() (*telebot.Bot, storage.Storage, parser.Parser) {
	token := mustToken()
	sqlitePath := mustSQLitePath()

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	strg, err := storage.NewSQLite(sqlitePath)
	if err != nil {
		log.Fatalf("error getting storage: %s", err.Error())
	}

	prsr := parser.NewGoq()

	return bot, strg, prsr
}

// mustToken получает токен из .env
func mustToken() string {
	return os.Getenv("TG_TOKEN")
}

// mustSQLitePath получает путь sqlite бд из .env
func mustSQLitePath() string {
	return os.Getenv("DBPATH")
}
