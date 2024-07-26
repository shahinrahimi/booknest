package store

import "github.com/shahinrahimi/booknest/pkg/book"

func (s *SqliteStore) GetBooks() ([]*book.Book, error) {
	rows, err := s.db.Query(book.SelectAll)
	if err != nil {
		return nil, err
	}
	var books []*book.Book
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b); err != nil {
			s.logger.Printf("Error scranning rows for a book: %v", err)
			continue
		}
		books = append(books, &b)
	}
	return books, nil
}

// GetBook return book and error if book not found it will return [ErrNoRows]
func (s *SqliteStore) GetBook(id string) (*book.Book, error) {
	row := s.db.QueryRow(book.Select, id)
	var b book.Book
	if err := row.Scan(b.ToFeilds()...); err != nil {
		return nil, err
	}
	return &b, nil
}

func (s *SqliteStore) CreateBook(b *book.Book) error {
	_, err := s.db.Exec(book.Insert, b.ToArgs()...)
	return err
}

func (s *SqliteStore) UpdateBook(id string, b *book.Book) error {
	_, err := s.db.Exec(book.Update, b.ToUpdatedArgs(id)...)
	return err
}

func (s *SqliteStore) DeleteBook(id string) error {
	_, err := s.db.Exec(book.Delete, id)
	return err
}
