package models

import "fmt"

// Class объект пары
type Class struct {
	Num     string
	Subject string
	Cabinet string
	Teacher string
	Notes   string
}

// String преобразует объект пары в строку
func (class *Class) String() string {
	// html теги для метода парсинга бота
	result := fmt.Sprintf(
		`<b>%s пара - %s</b>
<i>%s
%s</i>

`,
		class.Num,
		class.Cabinet,
		class.Subject,
		class.Teacher,
	)

	return result
}
