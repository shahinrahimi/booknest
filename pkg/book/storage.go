package book

type Storage interface {
	GetBooks() []*Book
	GetBook(id string) *Book
	CreateBook() error
	UpdateBook() error
	DeleteBook() error
}
