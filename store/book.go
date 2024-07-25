package store

import "github.com/shahinrahimi/booknest/pkg/book"

func (s *SqliteStore) GetBooks() ([]*book.Book, error) {
	return nil, nil
}

func (s *SqliteStore) GetBook(id string) (*book.Book, error) {
	return nil, nil
}

func (s *SqliteStore) CreateBook(b book.Book) error {
	return nil
}

func (s *SqliteStore) UpdateBook(id string, b book.Book) error {
	return nil
}

func (s *SqliteStore) DeleteBook(id string) error {
	return nil
}
