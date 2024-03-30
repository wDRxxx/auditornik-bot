package main

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
	"gopkg.in/telebot.v3"
)

// routes инициализирует роуты бота
func routes(a *config.AppConfig) {
	a.Bot.Handle("/start", handlers.Repo.Start)
	a.Bot.Handle("/help", handlers.Repo.Help)
	a.Bot.Handle("/about", handlers.Repo.About)

	a.Bot.Handle("/setgroup", handlers.Repo.SetGroup)

	a.Bot.Handle("/subscribe", handlers.Repo.UpdateMailing)
	a.Bot.Handle("/unsubscribe", handlers.Repo.UpdateMailing)

	a.Bot.Handle("/check", handlers.Repo.Check)
	a.Bot.Handle("/today", handlers.Repo.ScheduleToday)
	a.Bot.Handle("/tomorrow", handlers.Repo.ScheduleTomorrow)

	a.Bot.Handle(handlers.Repo.Keyboard.BtnToday, handlers.Repo.ScheduleToday)
	a.Bot.Handle(handlers.Repo.Keyboard.BtnTomorrow, handlers.Repo.ScheduleTomorrow)

	// обработка мусорных сообщений - ВСЕГДА В КОНЦЕ
	a.Bot.Handle(telebot.OnText, handlers.Repo.Unknown)
}
