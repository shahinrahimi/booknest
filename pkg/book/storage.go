package book

type Storage interface {
	GetBooks() ([]*Book, error)
	GetBook(id string) (*Book, error)
	CreateBook(b *Book) error
	UpdateBook(id string, b *Book) error
	DeleteBook(id string) error
}
