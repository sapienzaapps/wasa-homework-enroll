package database

func (db *appdbimpl) CreateStudent(studentId int, firstName string, lastName string, email string, repoURL string, publicKey string, privateKey string) error {
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	var cnt int
	err = tx.QueryRow(`SELECT COUNT(*) FROM students WHERE id=?`, studentId).Scan(&cnt)
	if err != nil {
		_ = tx.Rollback()
		return err
	} else if cnt > 0 {
		_ = tx.Rollback()
		return ErrUserExists
	}

	_, err = tx.Exec(`INSERT INTO students (id, first_name, last_name, email, repo_url, public_key, private_key) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		studentId, firstName, lastName, email, repoURL, publicKey, privateKey)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
