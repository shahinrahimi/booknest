package book

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Book defines the structure for an API book
// swagger: model Book
type Book struct {
	// ID of a book
	// required: false
	// in: path
	ID string `json:"id"`
	// the title for the book
	// example: Animal Farm
	// required: true
	Title string `json:"title" validate:"required"`
	// the author for the book
	// exmaple: George Orvil
	// required: true
	Author string `json:"author" validate:"required"`
	// the description for the book
	// example: The story about animals that lives in a farm
	// required: true
	Description string `json:"description" validate:"required"`
	// the cover for the book
	// example: http://...
	// required: true
	Cover string `json:"cover" validate:"required"`
	// the price for the book
	// example: 10
	// required: true
	// min: 1
	Price float32 `json:"price" validate:"required,gte=1"`
	// swagger: ignore
	CreatedAt time.Time `json:"-"`
	// swagger: ignore
	UpdatedAt time.Time `json:"-"`
}

type KeyBook struct{}

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT NOT NULL,
		cover TEXT NOT NULL,
		price REAL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);`
	SelectAll string = `SELECT id, title, author, description, cover, price, created_at FROM books`
	Select    string = `SELECT id, title, author, description, cover, price, created_at FROM books WHERE id = ?`
	Insert    string = `INSERT INTO books (id, title, author, description, cover, price, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	Update    string = `UPDATE books SET title = ?, author = ?, description = ?, cover = ?, price = ?, uptaded_at = ? WHERE id = ?`
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
		UpdatedAt:   time.Now().UTC(),
	}
}

// ToArgs returns id, description, author, cover, price, ceated_at, updated_at as value
func (b *Book) ToArgs() []interface{} {
	return []interface{}{b.ID, b.Title, b.Description, b.Author, b.Cover, b.Price, b.CreatedAt, b.UpdatedAt}
}

// ToUpdatedArgs returns title, description, author, cover, price, updated_at and id as value
func (b *Book) ToUpdatedArgs(id string) []interface{} {
	return []interface{}{b.Title, b.Description, b.Author, b.Cover, b.Price, b.UpdatedAt, id}
}

// ToFeilds returns id, title, description, author, cover, price, created_at, updated_at as refrence
func (b *Book) ToFeilds() []interface{} {
	return []interface{}{&b.ID, &b.Title, &b.Description, &b.Author, &b.Cover, &b.Price, &b.CreatedAt, &b.UpdatedAt}
}
