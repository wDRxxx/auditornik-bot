package handlers

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/storage"
	"gopkg.in/telebot.v3"
)

var Repo *Repository

type Repository struct {
	App     *config.AppConfig
	Storage storage.Storage
}

// NewRepository создает новый репозиторий
func NewRepository(a *config.AppConfig, s storage.Storage) *Repository {
	initGroups()
	return &Repository{
		App:     a,
		Storage: s,
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
// SetGroup обрабатывает выбор группы
func (m *Repository) SetGroup(c telebot.Context) error {
	tags := c.Args()
	if len(tags) == 0 {
		return c.Send(msgChooseGroupNoTag)
	}

	group := tags[0]

	_, exists := groups[group]
	if !exists {
		return c.Send(msgChooseGroupFail)
	}

	// TODO добавление записи в хранилище: юзер и его группы
	m.Storage.ZXC()

	return c.Send(msgChooseGroupSuccess + group)
}
