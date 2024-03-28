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
		`<b>%s пара</b>
<i>%s
%s</i>
<b>%s</b>

`,
		class.Num,
		class.Subject,
		class.Teacher,
		class.Cabinet,
	)

	return result
}
