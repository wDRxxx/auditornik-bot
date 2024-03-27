package parser

import "github.com/wDRxxx/auditornik-bot/internal/parser/goq"

type Parser interface {
	ScheduleForGroupByDate(groupId int, date string) (string, error)
}

func NewGoq() Parser {
	return &goq.Goq{}
}
