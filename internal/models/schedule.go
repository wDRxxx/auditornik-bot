package models

import "fmt"

// ScheduleDay объект дня расписания
type ScheduleDay struct {
	Event   string
	Classes []Class
}

// String преобразует объект дня расписания в строку
func (s *ScheduleDay) String() string {
	// на случай чего-то типо практики, указываем в самом начале
	var result = fmt.Sprintf("<b><i>%s</i></b>", s.Event)

	for _, class := range s.Classes {
		result += class.String()
	}

	return result
}
