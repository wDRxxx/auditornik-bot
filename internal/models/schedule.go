package models

import "fmt"

type ScheduleDay struct {
	Classes []Class
}

func (s *ScheduleDay) String() string {
	// TODO форматирование для ТГ
	var result string
	for _, class := range s.Classes {
		sclass := fmt.Sprintf(
			`%s пара
%s
%s
%s

`,
			class.Num,
			class.Subject,
			class.Teacher,
			class.Cabinet,
		)

		result += sclass
	}

	return result
}
