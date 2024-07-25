package store

import "github.com/shahinrahimi/booknest/pkg/book"

func (s *SqliteStore) GetBooks() []*book.Book {
	return nil
}

func (s *SqliteStore) GetBook(id string) *book.Book {
	return nil
}

func (s *SqliteStore) CreateBook() error {
	return nil
}

func (s *SqliteStore) UpdateBook() error {
	return nil
}

func (s *SqliteStore) DeleteBook() error {
	return nil
}
