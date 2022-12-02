package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

func (db *appdbimpl) GetDockerLog(studentid int) (string, error) {
	var log string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.c.QueryRow(ctx, `SELECT dockerlog FROM grades WHERE id=$1`, studentid).Scan(&log)
	cancel()
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrUserDoesNotExists
	}
	return log, err
}
