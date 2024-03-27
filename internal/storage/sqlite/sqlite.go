package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/wDRxxx/auditornik-bot/internal/helpers"
	"time"
)

type SQLite struct {
	DB *sql.DB
}

// UserGroup возвращает группу пользователя по его id
func (m *SQLite) UserGroup(userId int64) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select group_id from user_groups where user_id = ?`

	var groupId int
	err := m.DB.QueryRowContext(ctx, query, userId).Scan(&groupId)
	if err != nil {
		return groupId, helpers.ServerError("error getting group", err)
	}

	return groupId, nil
}

// SaveUserGroup сохраняет в хранилище пользователя и его группу
func (m *SQLite) SaveUserGroup(userId int64, groupId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lGroupId, err := m.UserGroup(userId)
	if lGroupId == groupId {
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		err = m.UpdateUserGroup(userId, groupId)
		return helpers.ServerError("error getting group", err)
	}

	query := `insert into user_groups (user_id, group_id) values (?, ?)`

	_, err = m.DB.ExecContext(ctx, query, userId, groupId)
	if err != nil {
		return helpers.ServerError("error saving user's group", err)
	}

	return nil
}

// UpdateUserGroup обновляет группу пользователя по его id
func (m *SQLite) UpdateUserGroup(userId int64, groupId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `update user_groups set group_id = ? where user_id = ?`

	_, err := m.DB.ExecContext(ctx, query, groupId, userId)
	if err != nil {
		return helpers.ServerError("error updating user's group", err)
	}

	return nil
}
