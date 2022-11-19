package database

import "context"

func (db *appdbimpl) CreateStudent(studentId int, firstName string, lastName string, email string, repoURL string, publicKey string, privateKey string) error {
	tx, err := db.c.Begin(context.Background())
	if err != nil {
		return err
	}

	var cnt int
	err = tx.QueryRow(context.Background(), `SELECT COUNT(*) FROM students WHERE id=$1 OR email=$2 OR repo_url=$3`, studentId, email, repoURL).Scan(&cnt)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	} else if cnt > 0 {
		_ = tx.Rollback(context.Background())
		return ErrUserExists
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO students (id, first_name, last_name, email, repo_url, public_key, private_key) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		studentId, firstName, lastName, email, repoURL, publicKey, privateKey)
	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}
	return tx.Commit(context.Background())
}
