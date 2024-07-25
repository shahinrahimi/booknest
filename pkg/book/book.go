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

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT NOT NULL,
		cover TEXT NOT NULL,
		price REAL,
		created_at TIMESTAMP NOT NULL
	);`
	SelectAll string = `SELECT id, title, author, description, cover, price, created_at FROM books`
	Select    string = `SELECT id, title, author, description, cover, price, created_at FROM books WHERE id = ?`
	Insert    string = `INSERT INTO books (id, title, author, description, cover, price, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	Update    string = `UPDATE books SET title = ?, author = ?, description = ?, cover = ?, price = ?, created_at = ? WHERE id = ?`
	Delete    string = `DELETE FROM books WHERE id = ?`
)

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
