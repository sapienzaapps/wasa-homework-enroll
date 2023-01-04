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
		var hr HomeworkResult
		err = rows.Scan(&hr.StudentID, &hr.Hash, &hr.OpenAPI, &hr.Go, &hr.Vue, &hr.Docker, &hr.LastCheck)
		if err != nil {
			return nil, err
		}

		ret = append(ret, hr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows.Close()

	return ret, nil
}
