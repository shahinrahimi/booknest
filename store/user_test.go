package store

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/shahinrahimi/booknest/pkg/user"
)

// Setup a test Logger
func setupTestLogger() *log.Logger {
	return log.New(os.Stdout, "[BOOKNEST-TEST] ", log.LstdFlags)
}

// Setup a test SQLite database
func setupTestDB(t *testing.T) *sql.DB {
	// Open sqlite db and try to connect
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	// Create users table for tests
	if _, err = db.Exec(user.CreateTable); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}
	return db
}

// Test GetUser
func TestGetUser(t *testing.T) {
	logger := setupTestLogger()
	db := setupTestDB(t)
	defer db.Close()
	store := &SqliteStore{db: db, logger: logger}

	// create a test user
	username := "testusername"
	password := "testpassword"
	newUser := user.NewUser(username, password)

	// insert user to DB
	if _, err := db.Exec(user.Insert, newUser.ToArgs()...); err != nil {
		t.Fatalf("failed to insert test user: %v", err)
	}

	// check if user is exists
	user, err := store.GetUser(newUser.ID)
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	// check if username is stored correctly
	if user.Username != newUser.Username {
		t.Fatalf("expected username '%s', got '%s'", newUser.Username, user.Username)
	}
}

func TestGetUsers(t *testing.T) {
	logger := setupTestLogger()
	db := setupTestDB(t)
	defer db.Close()
	store := &SqliteStore{db: db, logger: logger}

	users := []*user.User{
		{
			Username: "testusername1",
			Password: "testpassword1",
		},
		{
			Username: "testusername1",
			Password: "testpassword1",
		},
		{
			Username: "testusername1",
			Password: "testpassword1",
		},
	}

	for _, u := range users {
		newUser := user.NewUser(u.Username, u.Password)
		if err := store.CreateUser(*newUser); err != nil {
			t.Fatalf("error creating user: %v", err)
		}
	}

	// create a test user
	fetchedUsers, err := store.GetUsers()
	if err != nil {
		t.Fatalf("error query user from DB: %v", err)
	}

	// check users length
	if len(fetchedUsers) != len(users) {
		t.Fatalf("expected users length '%d', got '%d'", len(users), len(fetchedUsers))
	}
}
