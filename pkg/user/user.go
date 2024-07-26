package user

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		is_admin BOOLEAN
	);`
	SelectAll        string = `SELECT id, username, password, created_at, is_admin FROM users`
	SelectByID       string = `SELECT id, username, password, created_at, is_admin FROM users WHERE id = ?`
	SelectByUsername string = `SELECT id, username, password, created_at, is_admin FROM users WHERE username = ?`
	Insert           string = `INSERT INTO users (id, username, password, created_at, is_admin) VALUES (?, ?, ?, ?, ?)`
	Update           string = `UPDATE users SET username = ?, password = ?, created_at = ?, is_admin = ? WHERE id = ?`
	Delete           string = `DELETE FROM users WHERE id = ?`
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"-" validate:"required"`
	CreatedAt time.Time `json:"-"`
	IsAdmin   bool      `json:"is_admin"`
}

type KeyUser struct{}

func NewUser(username, password string) *User {
	hp, err := hashPassword(password)
	if err != nil {
		panic("Unable to hash the password")
	}
	return &User{
		ID:        fmt.Sprint("BU" + strconv.Itoa(rand.Int())),
		Username:  username,
		Password:  hp,
		CreatedAt: time.Now().UTC(),
		IsAdmin:   false,
	}
}

func NewRootUser(username, password string) *User {
	hp, err := hashPassword(password)
	if err != nil {
		panic("Unable to hash the password")
	}
	return &User{
		ID:        fmt.Sprint("BA" + strconv.Itoa(math.MaxInt)),
		Username:  username,
		Password:  hp,
		CreatedAt: time.Now().UTC(),
		IsAdmin:   true,
	}
}

func (u *User) ToArgs() []interface{} {
	return []interface{}{u.ID, u.Username, u.Password, u.CreatedAt, u.IsAdmin}
}
func (u *User) ToUpdatedArgs() []interface{} {
	return []interface{}{u.Username, u.Password, u.CreatedAt, u.IsAdmin, u.ID}
}
func (u *User) ToFeilds() []interface{} {
	return []interface{}{&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.IsAdmin}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
