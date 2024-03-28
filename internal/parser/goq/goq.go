package goq

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wDRxxx/auditornik-bot/internal/helpers"
	"github.com/wDRxxx/auditornik-bot/internal/models"
	"log"
	"net/http"
	"strings"
)

type Goq struct{}

// ScheduleForGroupByDate возвращает расписание для группы
func (m *Goq) ScheduleForGroupByDate(groupId int, date string) (string, error) {
	url := fmt.Sprintf("https://ies.unitech-mo.ru/schedule_list_groups?d=%s&g=%d", date, groupId)
	doc, err := document(url)
	if err != nil {
		return "", helpers.ServerError("error getting schedule for group", err)
	}

	schedule, err := parseSchedule(doc, date)
	if err != nil {
		return "", err
	}

	return schedule.String(), nil
}

// parseSchedule получает расписание на день из документа
func parseSchedule(doc *goquery.Document, date string) (models.ScheduleDay, error) {
	// содержит объект таблицы с расписанием на выбранный день
	var scheduleTable *goquery.Selection
	doc.Find(".text-center").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), date) {
			scheduleTable = s.Next().Next()
			return
		}
	})

	if scheduleTable == nil {
		return models.ScheduleDay{}, errors.New("error getting schedule table")
	}

	// получает данные расписания на определенный день
	var schedule models.ScheduleDay
	scheduleTable.Find("tbody").Find("tr").Each(func(_ int, sel *goquery.Selection) {
		var class models.Class
		sel.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				class.Num = strings.TrimSpace(s.Text())
			case 2:
				exploded := strings.Split(strings.TrimSpace(s.Text()), "-")
				log.Println(exploded)
				class.Subject = exploded[0]
				if class.Subject == "" {
					return
				}
			case 3:
				class.Cabinet = strings.TrimSpace(s.Text())
			case 4:
				class.Teacher = strings.TrimSpace(s.Text())
			case 5:
				class.Notes = strings.TrimSpace(s.Text())
			}
		})

		if class.Subject != "" {
			schedule.Classes = append(schedule.Classes, class)
		}
	})

	return schedule, nil
}

// document возвращает объект документа по url для дальнейшего парсинга.
func document(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return nil, helpers.ServerError("error getting body from url", err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)

	return doc, nil
}
