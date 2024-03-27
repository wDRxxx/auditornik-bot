package handlers

import (
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/parser"
	"github.com/wDRxxx/auditornik-bot/internal/storage"
	"gopkg.in/telebot.v3"
	"strings"
	"time"
)

var Repo *Repository

type Repository struct {
	App     *config.AppConfig
	Storage storage.Storage
	Parser  parser.Parser
}

// NewRepository создает новый репозиторий
func NewRepository(a *config.AppConfig, s storage.Storage, p parser.Parser) *Repository {
	initGroups()
	return &Repository{
		App:     a,
		Storage: s,
		Parser:  p,
	}
}

// NewHandlers обновляет Repo
func NewHandlers(r *Repository) {
	Repo = r
}

// Unknown обрабатывае мусорные сообщения
func (m *Repository) Unknown(c telebot.Context) error {
	return c.Send(msgUnknown)
}

// Start отправляет приветственное сообщение
func (m *Repository) Start(c telebot.Context) error {
	return c.Send(msgStart)
}

// SetGroup обрабатывае и отправляет запрос к бд для сохранения группы у пользователя
func (m *Repository) SetGroup(c telebot.Context) error {
	tags := c.Args()
	if len(tags) == 0 {
		return c.Send(msgChooseGroupNoTag)
	}

	groupName := tags[0]

	groupId, exists := groups[strings.ToUpper(groupName)]
	if !exists {
		return c.Send(msgChooseGroupFail)
	}

	err := m.Storage.SaveUserGroup(c.Sender().ID, groupId)
	if err != nil {
		return err
	}

	return c.Send(msgChooseGroupSuccess + groupName)
}

// ScheduleToday отправляет расписание на сегодня
func (m *Repository) ScheduleToday(c telebot.Context) error {
	// отправляет сообщение через объект бота, сохраняя его.
	// если отправлять через контекст - нельзя получить объект сообщения
	// и отредактировать его
	msg, err := m.App.Bot.Send(c.Chat(), "Секунду...")
	if err != nil {
		return err
	}

	date := time.Now()
	today := date.Format("02.01.2006")

	groupId, err := m.Storage.UserGroup(c.Sender().ID)
	if err != nil {
		return err
	}

	schedule, err := m.Parser.ScheduleForGroupByDate(groupId, today)
	if err != nil {
		return err
	}

	msg, err = m.App.Bot.Edit(msg, schedule)
	if err != nil {
		return err
	}

	return nil
}

// ScheduleTomorrow отправляет расписание на завтра
func (m *Repository) ScheduleTomorrow(c telebot.Context) error {
	return c.Send("Завтра")
}

// ScheduleDATomorrow отправляет расписание на послезавтра
func (m *Repository) ScheduleDATomorrow(c telebot.Context) error {
	return c.Send("Послезавтра")
}
