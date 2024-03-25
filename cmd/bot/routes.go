package main

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
)

func routes(a *config.AppConfig) {
	a.Bot.Handle("/start", handlers.Repo.Start)
}
