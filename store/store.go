package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shahinrahimi/booknest/pkg/book"
	"github.com/shahinrahimi/booknest/pkg/user"
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

func (s *SqliteStore) Init() {
	// create users table if not exists
	if _, err := s.db.Exec(user.CreateTable); err != nil {
		s.logger.Printf("Error creating users table: %v", err)
		s.logger.Panic("Unable to create users table")
	}

	// create books table if not exists
	if _, err := s.db.Exec(book.CreateTable); err != nil {
		s.logger.Printf("Error creating books table: %v", err)
		s.logger.Panic("Unable to create books table")
	}
}

func (s *SqliteStore) Close() error {
	return s.db.Close()
}
