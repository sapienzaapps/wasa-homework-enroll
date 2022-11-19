/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserExists        = errors.New("user exists")
	ErrUserDoesNotExists = errors.New("user does not exists")
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// CreateStudent creates a new user if he/she doesn't exist
	CreateStudent(studentId int, firstName string, lastName string, email string, repoURL string, publicKey string, privateKey string) error

	ListResults() ([]HomeworkResult, error)
	GetGitLog(studentid int) (string, error)
	GetOpenAPILog(studentid int) (string, error)

	Ping() error
}

type appdbimpl struct {
	c *pgxpool.Pool
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *pgxpool.Pool) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping(context.Background())
}
