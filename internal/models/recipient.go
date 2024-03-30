package models

// Recipient структура для работы с telebot.Recipient
type Recipient struct {
	Id string
}

// Recipient возвращает id чата в формате строки
func (r Recipient) Recipient() string {
	return r.Id
}

//msg, err := m.App.Bot.Send(&Recipient{Id: "826582348"}, msgWait)
