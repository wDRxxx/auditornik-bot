package config

import (
	"github.com/wDRxxx/auditornik-bot/pkg/job-ticker"
	"gopkg.in/telebot.v3"
)

// AppConfig содержит конфигурацию приложения
type AppConfig struct {
	Bot    *telebot.Bot
	Ticker *job_ticker.JobTicker
}
