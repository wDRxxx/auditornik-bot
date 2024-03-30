package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/wDRxxx/auditornik-bot/internal/helpers"
	"github.com/wDRxxx/auditornik-bot/internal/models"
	"time"
)

const (
	errStrUpdatingGroup       = "error updating user's group"
	errStrSavingGroup         = "error saving user's group"
	errStrGettingGroup        = "error getting user's group"
	errStrUpdatingMailing     = "error updating user's mailing"
	errStrGettingUsersMailing = "error getting user's mailing"
)

type SQLite struct {
	DB *sql.DB
}

// UserGroup возвращает группу пользователя
func (m *SQLite) UserGroup(userId int64) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select group_id from user_groups where user_id = ?`

	var groupId int
	err := m.DB.QueryRowContext(ctx, query, userId).Scan(&groupId)
	if errors.Is(err, sql.ErrNoRows) {
		return groupId, err
	}
	if err != nil {
		return groupId, helpers.ServerError(errStrGettingGroup, err)
	}

	return groupId, nil
}

// SaveUserGroup сохраняет в хранилище пользователя и его группу
func (m *SQLite) SaveUserGroup(userId int64, username string, groupId int) error {
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

	query := `insert into user_groups (user_id, username, group_id) values (?, ?, ?)`

	_, err = m.DB.ExecContext(ctx, query, userId, username, groupId)
	if err != nil {
		return helpers.ServerError(errStrSavingGroup, err)
	}

	return nil
}

// UpdateUserGroup обновляет группу пользователя
func (m *SQLite) UpdateUserGroup(userId int64, groupId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `update user_groups set group_id = ? where user_id = ?`

	_, err := m.DB.ExecContext(ctx, query, groupId, userId)
	if err != nil {
		return helpers.ServerError(errStrUpdatingGroup, err)
	}

	return nil
}

// UpdateUserMailing обновляет состояние подписки на рассылку пользователя
func (m *SQLite) UpdateUserMailing(userId int64, mailingStatus int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `update user_groups set mailing = ? where user_id = ?`

	_, err := m.DB.ExecContext(ctx, query, mailingStatus, userId)
	if err != nil {
		return helpers.ServerError(errStrUpdatingMailing, err)
	}

	return nil
}

// AllUsersWithMailing получает и возвращает всех пользователей с подпиской на рассылку
func (m *SQLite) AllUsersWithMailing() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	query := `select user_id, username, group_id  from user_groups where mailing = ?`

	rows, err := m.DB.QueryContext(ctx, query, 1)
	defer rows.Close()

	if err != nil {
		return users, helpers.ServerError(errStrGettingUsersMailing, err)
	}

	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.GroupId,
		)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}
