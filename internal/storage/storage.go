package storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wDRxxx/auditornik-bot/internal/models"
	"github.com/wDRxxx/auditornik-bot/internal/storage/sqlite"
)

var (
	ErrNoClasses       = errors.New("no lessons in chosen day")
	ErrGettingSchedule = errors.New("error getting schedule table")
)

type Storage interface {
	UserGroup(userId int64) (int, error)
	SaveUserGroup(userId int64, username string, groupId int) error
	UpdateUserGroup(userId int64, groupId int) error
	UpdateUserMailing(userId int64, mailingStatus int) error
	AllUsersWithMailing() ([]models.User, error)
}

func NewSQLite(path string) (Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &sqlite.SQLite{DB: db}, nil
}
