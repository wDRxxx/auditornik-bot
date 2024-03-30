package job_ticker

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/wDRxxx/auditornik-bot/internal/helpers"
	"log"
	"time"
)

const HourToTick uint = 19
const MinuteToTick uint = 00

// JobTicker - тикер для периодического запуска необходимых функций
type JobTicker struct {
	S gocron.Scheduler
}

// New возвращает новый тикер
func New(f func() error) *JobTicker {
	s, _ := gocron.NewScheduler()

	_, _ = s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(HourToTick, MinuteToTick, 0),
			),
		),
		gocron.NewTask(
			func() {
				if time.Now().Weekday() != 6 {
					err := f()
					if err != nil {
						log.Println(helpers.ServerError("error doing gocron task", err))
					}
				}
			},
		),
	)

	return &JobTicker{S: s}
}

// Run запускает ежедневную рассылку (кроме воскресенья и дней когда пар нет)
func (jt *JobTicker) Run() {
	defer jt.S.Shutdown()
	jt.S.Start()

	select {}
}
