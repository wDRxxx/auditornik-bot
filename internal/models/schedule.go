package models

// ScheduleDay объект дня расписания
type ScheduleDay struct {
	Classes []Class
}

// String преобразует объект дня расписания в строку
func (s *ScheduleDay) String() string {
	// TODO форматирование для ТГ
	var result string
	for _, class := range s.Classes {
		result += class.String()
	}

	return result
}
