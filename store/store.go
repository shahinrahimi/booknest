package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	logger *log.Logger
	db     *sql.DB
}

func NewSqliteStore(logger *log.Logger) *SqliteStore {
	// create directory if not exists
	if err := os.MkdirAll("db", 0755); err != nil {
		logger.Panic("Unable to create a directory for DB")
	}

	// create connection to db
	db, err := sql.Open("sqlite3", "./db/booknest.db")
	if err != nil {
		logger.Panic("Unable to connect to DB")
	}
	logger.Println("DB Connected!")

	return &SqliteStore{
		logger: logger,
		db:     db,
	}
}

func (s *SqliteStore) Close() error {
	return s.db.Close()
}
