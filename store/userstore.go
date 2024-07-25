package store

import "github.com/shahinrahimi/booknest/pkg/user"

func (s *SqliteStore) GetUsers() []*user.User {
	return nil
}

func (s *SqliteStore) GetUser(id string) *user.User {
	return nil
}

func (s *SqliteStore) CreateUser() error {
	return nil
}

func (s *SqliteStore) UpdateUser() error {
	return nil
}

func (s *SqliteStore) DeleteUser() error {
	return nil
}
