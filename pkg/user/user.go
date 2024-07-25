package user

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func NewUser(username, password string) *User {
	hp, err := hashPassword(password)
	if err != nil {
		panic("Unable to hash the password")

	}
	return &User{
		ID:        fmt.Sprint("BB" + strconv.Itoa(rand.Int())),
		Username:  username,
		Password:  hp,
		CreatedAt: time.Now().UTC(),
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
