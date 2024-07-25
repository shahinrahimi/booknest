package book

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Cover       string    `json:"cover"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"-"`
}

func NewBook(title, author, description, cover string, price float32) *Book {
	return &Book{
		ID:          fmt.Sprint("BB" + strconv.Itoa(rand.Int())),
		Title:       title,
		Author:      author,
		Description: description,
		Cover:       cover,
		Price:       price,
		CreatedAt:   time.Now().UTC(),
	}
}
