package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (db *appdbimpl) GetGitLog(studentid int) (string, error) {
	var log string
	err := db.c.QueryRow(context.Background(), `SELECT gitlog FROM grades WHERE id=$1`, studentid).Scan(&log)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrUserDoesNotExists
	}
	return log, err
}
