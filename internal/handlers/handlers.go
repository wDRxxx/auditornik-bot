package handlers

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"gopkg.in/telebot.v3"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository создает новый репозиторий
func NewRepository(a *config.AppConfig) *Repository {
	initGroups()
	return &Repository{
		App: a,
	}
}

// NewHandlers обновляет Repo
func NewHandlers(r *Repository) {
	Repo = r
}

// Start отправляет прветсвтенное сообщение
func (m *Repository) Start(c telebot.Context) error {
	return c.Send(msgStart)
}

// ChooseGroup обрабатывает выбор группы
func (m *Repository) ChooseGroup(c telebot.Context) error {
	tags := c.Args()
	if len(tags) == 0 {
		return c.Send(msgChooseGroupNoTag)
	}

	group := tags[0]

	groupId, exists := groups[group]
	if !exists {
		return c.Send(msgChooseGroupFail)
	}

	// TODO добавление записи в хранилище: юзер и его группы

	return c.Send(msgChooseGroupSuccess + group)
}
