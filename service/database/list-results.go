package database

import (
	"context"
	"time"
)

type HomeworkResult struct {
	StudentID uint64
	Hash      string
	OpenAPI   int
	Go        int
	Vue       int
	Docker    int
	LastCheck time.Time
}

func (db *appdbimpl) ListResults() ([]HomeworkResult, error) {
	var ret []HomeworkResult

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	rows, err := db.c.Query(ctx, `SELECT id, hash, openapi, golang, vue, docker, lastcheck FROM grades`)
	cancel()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var hr HomeworkResult
		err = rows.Scan(&hr.StudentID, &hr.Hash, &hr.OpenAPI, &hr.Go, &hr.Vue, &hr.Docker, &hr.LastCheck)
		if err != nil {
			return nil, err
		}

		ret = append(ret, hr)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	rows.Close()

	return ret, nil
}
