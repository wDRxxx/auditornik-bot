package main

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/handlers"
	"gopkg.in/telebot.v3"
)

// routes инициализирует роуты бота
func routes(a *config.AppConfig) {
	a.Bot.Handle("/start", handlers.Repo.Start)
	a.Bot.Handle("/setgroup", handlers.Repo.SetGroup)

	// TODO переделать получение расписания с обычных комманд на кнопки клавиатуры ТГ
	a.Bot.Handle("/curr", handlers.Repo.ScheduleToday)
	a.Bot.Handle("/next", handlers.Repo.ScheduleTomorrow)
	a.Bot.Handle("/next2", handlers.Repo.ScheduleDATomorrow)

	// обработка мусорных сообщений - ВСЕГДА В КОНЦЕ
	a.Bot.Handle(telebot.OnText, handlers.Repo.Unknown)
}
