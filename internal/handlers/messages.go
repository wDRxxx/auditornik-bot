package handlers

const (
	msgUnknown = `Не знаю такого`
	msgStart   = `Привет! Я бот-аудиторник для портала ies.unitech-mo.ru.
Для выбора группы напиши /setgroup <group>.
Пример: /setgroup П1-22

Список команд: /help`

	msgChooseGroupSuccess = `Запомнил твою группу: %s.
Для получения расписания выбери день в клавиатуре тг или напиши /today /tomorrow
Для подписки на рассылку расписания напиши /subscribe
`

	msgChooseGroupNoTag = `Ты не указал группу, попробуй снова.`
	msgChooseGroupFail  = `Не нашел такой группы`
	msgNoGroup          = `Что-то не так, может ты забыл указать группу?`
	msgWrongDay         = `Что-то не так, ты неправильно указал день для расписния.
Пример: /check П1-22 завтра
`

	msgSubscribe = `Теперь ты будешь получать расписание на следующий день в 19:00 ежедневно!
Чтобы отписаться от рассылки напиши /unsubscribe`
	msgUnsubscribe = `Теперь ты не будешь получать расписание.`

	msgWait      = `Секунду...`
	msgNoClasses = `В этот день пар нет!`

	msgHelp = `Список команд:
/setgroup <group> - Выбор группы для получения расписания;
/today или "Расписание на сегодня" на клавиатуре тг - Получение расписания на сегодня;
/tomorrow или "Расписание на завтра" на клавиатуре тг - Получение расписания на завтра;
/check <П1-22> <сегодня\завтра> - Получение расписания для указанной группы;
/subscribe - Включение ежедневной рассылки расписания;
/unsubscribe - Выключение ежедневной рассылки расписания;
/about - доп. техническая информация.
`
	msgAbout = `tg: @drx_links

Бот написан на go v1.21;
Исходный код - https://github.com/wDRxxx/auditornik-bot
`
)
