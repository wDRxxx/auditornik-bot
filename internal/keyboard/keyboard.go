package keyboard

import "gopkg.in/telebot.v3"

// Keyboard содержит меню и его кнопки
type Keyboard struct {
	Menu        *telebot.ReplyMarkup
	BtnToday    *telebot.Btn
	BtnTomorrow *telebot.Btn
}

// New создает новый объект (неожиданно)
func New() *Keyboard {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnToday := menu.Text("Расписание на сегодня")
	btnTomorrow := menu.Text("Расписание на завтра")

	menu.Reply(
		menu.Row(btnToday),
		menu.Row(btnTomorrow),
	)

	return &Keyboard{
		Menu:        menu,
		BtnToday:    &btnToday,
		BtnTomorrow: &btnTomorrow,
	}
}
