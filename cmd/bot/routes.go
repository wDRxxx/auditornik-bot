package main

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
)

// routes инициализирует роуты бота
func routes(a *config.AppConfig) {
	a.Bot.Handle("/start", handlers.Repo.Start)
	a.Bot.Handle("/choose", handlers.Repo.ChooseGroup)
}
