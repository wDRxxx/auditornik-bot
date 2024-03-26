package sqlite

import (
	"database/sql"
	"fmt"
)

type SQLite struct {
	DB *sql.DB
}

func (s *SQLite) ZXC() {
	fmt.Println("SQLite")
}
