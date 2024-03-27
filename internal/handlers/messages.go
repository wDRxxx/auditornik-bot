package handlers

const (
	msgUnknown = `Не знаю такого`
	msgStart   = `Привет! Я бот-аудиторник для портала ies.unitech-mo.ru.
Для выбора группы введи /setgroup <group>.
Пример: /setgroup П1-22`
	msgChooseGroupNoTag = `Ты не указал группу, попробуй снова.
Пример: /setgroup П1-22`
	// TODO улучшить msgChooseGroupSuccess
	msgChooseGroupSuccess = `Запомнил твою группу: `
	msgChooseGroupFail    = `Не нашел такой группы`
)
