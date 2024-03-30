package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/wDRxxx/auditornik-bot/internal/config"
	"github.com/wDRxxx/auditornik-bot/internal/keyboard"
	"github.com/wDRxxx/auditornik-bot/internal/models"
	"github.com/wDRxxx/auditornik-bot/internal/parser"
	"github.com/wDRxxx/auditornik-bot/internal/storage"
	job_ticker "github.com/wDRxxx/auditornik-bot/pkg/job-ticker"
	"gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Repo *Repository

type Repository struct {
	App      *config.AppConfig
	Storage  storage.Storage
	Parser   parser.Parser
	Keyboard *keyboard.Keyboard
	Ticker   *job_ticker.JobTicker
}

// NewRepository создает новый репозиторий
func NewRepository(a *config.AppConfig, s storage.Storage, p parser.Parser) *Repository {
	initGroups()
	menu := keyboard.New()

	return &Repository{
		App:      a,
		Storage:  s,
		Parser:   p,
		Keyboard: menu,
	}
}

// NewHandlers обновляет Repo
func NewHandlers(r *Repository) {
	Repo = r
}

// Mailing делает рассылку всем пользователям подписанным на нее
func (m *Repository) Mailing() error {
	// создается мапа расписания по группам, если группы нет в мапе -
	// создем расписание и кладем туда, иначе просто берем из мапы
	schedules := make(map[int]string)
	users, err := m.Storage.AllUsersWithMailing()
	if err != nil {
		return err
	}

	tomorrow := time.Now().AddDate(0, 0, 1).Format("02.01.2006")

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func(wg *sync.WaitGroup, user models.User) {
			defer wg.Done()
			log.Println(user)

			schedule, exists := schedules[user.GroupId]
			if exists {
				err = m.sendMailingSchedule(schedule, user.Id)
				if err != nil {
					log.Println(err)
				}
			}

			scheduleString, err := m.Parser.ScheduleForGroupByDate(user.GroupId, tomorrow)
			schedules[user.GroupId] = scheduleString

			if errors.Is(err, storage.ErrNoClasses) {
				return
			}
			if err != nil {
				log.Println(err)
				return
			}

			err = m.sendMailingSchedule(scheduleString, user.Id)

			if err != nil {
				log.Println(err)
				return
			}
		}(&wg, user)
	}

	wg.Wait()
	log.Println("Mailing was successes!")

	return nil
}

// Unknown обрабатывае мусорные сообщения
func (m *Repository) Unknown(c telebot.Context) error {
	return c.Send(msgUnknown)
}

// Start отправляет приветственное сообщение
func (m *Repository) Start(c telebot.Context) error {
	return c.Send(msgStart)
}

// Help отправляет help сообщение
func (m *Repository) Help(c telebot.Context) error {
	return c.Send(msgHelp)
}

// About отправляет about сообщение
func (m *Repository) About(c telebot.Context) error {
	return c.Send(msgAbout)
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

	err := m.Storage.SaveUserGroup(c.Sender().ID, c.Sender().Username, groupId)
	if err != nil {
		return err
	}

	return c.Send(fmt.Sprintf(msgChooseGroupSuccess, strings.ToUpper(groupName)), Repo.Keyboard.Menu)
}

func (m *Repository) UpdateMailing(c telebot.Context) error {
	if c.Message().Text == "/subscribe" {
		err := m.Storage.UpdateUserMailing(c.Sender().ID, 1)
		if err != nil {
			return err
		}
		return c.Send(msgSubscribe)
	}

	err := m.Storage.UpdateUserMailing(c.Sender().ID, 0)
	if err != nil {
		return err
	}
	return c.Send(msgUnsubscribe)
}

// ScheduleToday отправляет расписание на сегодня
func (m *Repository) ScheduleToday(c telebot.Context) error {
	msg, err := m.sendWaitMessage(c)
	if err != nil {
		return err
	}

	groupId, err := m.userGroup(msg, c)
	if err != nil {
		return err
	}

	date := time.Now()
	today := date.Format("02.01.2006")

	return m.sendScheduleForGroupByDate(msg, groupId, today, c)
}

// ScheduleTomorrow отправляет расписание на завтра
func (m *Repository) ScheduleTomorrow(c telebot.Context) error {
	msg, err := m.sendWaitMessage(c)
	if err != nil {
		return err
	}

	groupId, err := m.userGroup(msg, c)
	if err != nil {
		return err
	}

	date := time.Now().AddDate(0, 0, 1)
	tomorrow := date.Format("02.01.2006")

	return m.sendScheduleForGroupByDate(msg, groupId, tomorrow, c)
}

// Check отправляет расписания для группы указанной в сообщении
func (m *Repository) Check(c telebot.Context) error {
	msg, err := m.sendWaitMessage(c)
	if err != nil {
		return err
	}

	tags := c.Args()
	if len(tags) == 0 {
		msg, err = m.App.Bot.Edit(msg, msgChooseGroupNoTag)
		return err
	}

	groupName := tags[0]
	groupId, exists := groups[strings.ToUpper(groupName)]
	if !exists {
		msg, err = m.App.Bot.Edit(msg, msgChooseGroupFail)
		return err
	}

	day := tags[1]

	date := time.Now()
	var dateString = "мусор"

	if strings.ToLower(day) == "сегодня" {
		dateString = date.Format("02.01.2006")
	} else if strings.ToLower(day) == "завтра" {
		dateString = date.AddDate(0, 0, 1).Format("02.01.2006")
	} else {
		msg, err = m.App.Bot.Edit(msg, msgWrongDay)
		return err
	}

	return m.sendScheduleForGroupByDate(msg, groupId, dateString, c)
}

// sendScheduleForGroupByDate отправляет расписание для группы пользователя на выбранную дату
func (m *Repository) sendScheduleForGroupByDate(msg *telebot.Message, groupId int, date string, c telebot.Context) error {
	schedule, err := m.Parser.ScheduleForGroupByDate(groupId, date)
	if errors.Is(err, storage.ErrNoClasses) {
		_, err := m.App.Bot.Edit(msg, msgNoClasses)
		return err
	}
	if err != nil {
		return err
	}

	// html парсинг для правильного отображения сообщения
	msg, err = m.App.Bot.Edit(msg, schedule, &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	})
	if err != nil {
		return err
	}

	return nil
}

// sendWaitMessage отправляет сообщение через объект бота, сохраняя его.
// Если же отправлять как обычно, через контекст -
// нельзя получить объект сообщения и отредактировать его.
func (m *Repository) sendWaitMessage(c telebot.Context) (*telebot.Message, error) {
	msg, err := m.App.Bot.Send(c.Chat(), msgWait)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

// userGroup получает айди группы пользователя
func (m *Repository) userGroup(msg *telebot.Message, c telebot.Context) (int, error) {
	groupId, err := m.Storage.UserGroup(c.Sender().ID)
	if errors.Is(err, sql.ErrNoRows) {
		_, err := m.App.Bot.Edit(msg, msgNoGroup)
		return groupId, err
	}
	if err != nil {
		return groupId, err
	}

	return groupId, nil
}

func (m *Repository) sendMailingSchedule(schedule string, userId int) error {
	schedule = "Расписание на завтра\n" + schedule

	r := models.Recipient{Id: strconv.Itoa(userId)}
	_, err := m.App.Bot.Send(r, schedule, &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	})

	return err
}
