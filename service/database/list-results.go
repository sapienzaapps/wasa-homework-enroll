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

	rows, err := db.c.Query(context.Background(), `SELECT id, hash, openapi, golang, vue, docker, lastcheck FROM grades`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}

		var hr HomeworkResult
		err = rows.Scan(&hr.StudentID, &hr.Hash, &hr.OpenAPI, &hr.Go, &hr.Vue, &hr.Docker, &hr.LastCheck)
		if err != nil {
			return nil, err
		}

		ret = append(ret, hr)
	}
	rows.Close()

	return ret, nil
}
