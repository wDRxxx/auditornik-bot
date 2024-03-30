# Auditornik-bot
Auditornik-bot - телеграм-бот для получения расписания с ies.unitech-mo.ru

Использовано:
- Go v1.21
- [Telebot](https://github.com/tucnak/telebot)
- [Goquery](https://github.com/PuerkitoBio/goquery)
- [Godotenv](https://github.com/joho/godotenv)
- [Go-sqlite3](https://github.com/mattn/go-sqlite3)
- [Gocron](https://github.com/go-co-op/gocron)

# Локальный запуск
- Создать sqlite бд и использовать scheme.sql;
- Переименовать .env-example на .env;
- Заполнить .env своими значениями;
- Выполнить ` go build ./cmd/bot`