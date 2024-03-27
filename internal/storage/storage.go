package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wDRxxx/auditornik-bot/internal/storage/sqlite"
)

type Storage interface {
	ZXC()
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