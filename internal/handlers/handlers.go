package handlers

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"gopkg.in/telebot.v3"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Start(c telebot.Context) error {
	return c.Send(msgStart)
}
